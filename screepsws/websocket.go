package screepsws

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/hinshun/screepsapi/screepstype"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	conn          *websocket.Conn
	serverURL     *url.URL
	token         string
	interrupt     chan struct{}
	authenticated bool
	authLock      sync.RWMutex
	sendQueue     []string
	subscriptions map[string]chan []byte
}

func NewWebSocket(serverURL *url.URL, token string) *WebSocket {
	return &WebSocket{
		serverURL:     serverURL,
		token:         token,
		interrupt:     make(chan struct{}),
		subscriptions: make(map[string]chan []byte),
	}
}

func (ws *WebSocket) Connect() error {
	if ws.conn != nil {
		return fmt.Errorf("websocket is already connected")
	}

	websocketURL, _ := url.Parse(ws.serverURL.String())
	websocketURL.Scheme = strings.Replace(websocketURL.Scheme, "http", "ws", 1)
	websocketURL.Path = socketPath

	conn, _, err := websocket.DefaultDialer.Dial(websocketURL.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create websocket connection: %s", err)
	}
	ws.conn = conn

	err = ws.Authenticate(ws.token)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %s", err)
	}

	err = ws.SetGZIP(true)
	if err != nil {
		return fmt.Errorf("failed to enable gzip: %s", err)
	}

	go ws.Listen()

	return nil
}

func (ws *WebSocket) Close() error {
	if ws.conn == nil {
		return fmt.Errorf("websocket is not connected")
	}

	close(ws.interrupt)
	ws.interrupt = make(chan struct{})
	ws.authLock.Lock()
	ws.authenticated = false
	ws.sendQueue = ws.sendQueue[:0]
	ws.authLock.Unlock()

	err := ws.conn.Close()
	if err != nil {
		return err
	}
	ws.conn = nil

	return nil
}

func (ws *WebSocket) Send(data string) error {
	ws.authLock.RLock()
	defer ws.authLock.RUnlock()
	if !ws.authenticated {
		ws.sendQueue = append(ws.sendQueue, data)
		return nil
	}

	fmt.Printf("websocket: %s\n", data)
	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		return fmt.Errorf("failed to send '%s': %s", data, err)
	}
	return nil
}

func (ws *WebSocket) Receive() (data []byte, err error) {
	_, data, err = ws.conn.ReadMessage()
	if err != nil {
		return
	}

	limit := len(data)
	suffix := ""
	if limit > 50 {
		limit = 50
		suffix = "..."
	}
	fmt.Printf("websocket-data: %s%s\n", data[:limit], suffix)
	return
}

func (ws *WebSocket) Authenticate(token string) error {
	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(authFormat, ws.token)))
	if err != nil {
		return fmt.Errorf("failed to authenticate: %s", err)
	}
	return nil
}

func (ws *WebSocket) SetGZIP(enable bool) error {
	arg := "off"
	if enable {
		arg = "on"
	}

	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(gzipFormat, arg)))
	if err != nil {
		return fmt.Errorf("failed to authenticate: %s", err)
	}
	return nil
}

func (ws *WebSocket) Subscribe(channel string) (<-chan []byte, error) {
	_, exists := ws.subscriptions[channel]
	if exists {
		return nil, fmt.Errorf("channel '%s' already subscribed", channel)
	}

	err := ws.Send(fmt.Sprintf(subscribeFormat, channel))
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe: %s", err)
	}

	dataChan := make(chan []byte)
	ws.subscriptions[channel] = dataChan

	return dataChan, nil
}

func (ws *WebSocket) Unsubscribe(channel string) error {
	err := ws.Send(fmt.Sprintf(unsubscribeFormat, channel))
	if err != nil {
		return fmt.Errorf("failed to unsubscribe: %s", err)
	}

	dataChan, exists := ws.subscriptions[channel]
	if exists {
		close(dataChan)
		delete(ws.subscriptions, channel)
	}

	return nil
}

func (ws *WebSocket) Listen() {
	for i := 0; i < 4; i++ {
		_, err := ws.Receive()
		if err != nil {
			fmt.Printf("failed to receive connection handshake: %s\n", err)
			return
		}
	}

	ws.authLock.Lock()
	ws.authenticated = true
	ws.authLock.Unlock()

	for len(ws.sendQueue) > 0 {
		data := ws.sendQueue[0]
		err := ws.Send(data)
		if err != nil {
			fmt.Printf("failed to send data off queue '%s': %s\n", data, err)
		}
		ws.sendQueue = ws.sendQueue[1:]
	}

	for {
		select {
		case <-ws.interrupt:
			return
		default:
			data, err := ws.Receive()
			if err != nil {
				fmt.Printf("failed to receive data: %s\n", err)
				continue
			}

			if len(data) < len(screepstype.GzipPrefix)+2 {
				continue
			}

			if string(data[:len(screepstype.GzipPrefix)]) == screepstype.GzipPrefix {
				err = ws.handleGzippedData(data)
				if err != nil {
					fmt.Printf("failed to handle gzipped data: %s\n", err)
				}
			} else {
				err = ws.handleData(data)
				if err != nil {
					fmt.Printf("failed to handle data: %s\n", err)
				}
			}
		}
	}
}

func (ws *WebSocket) handleData(data []byte) error {
	resp := make([]json.RawMessage, 2)
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal received data '%s': %s\n", data, err)
	}

	channel, err := resp[0].MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal channel name: %s", err)
	}

	channelData, err := resp[1].MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal channel data: %s", err)
	}

	subscription, ok := ws.subscriptions[string(channel[1:len(channel)-1])]
	if ok {
		subscription <- channelData
	}

	return nil
}

func (ws *WebSocket) handleGzippedData(data []byte) error {
	unzippedData, err := screepstype.Unzip(string(data))
	if err != nil {
		return fmt.Errorf("failed to unzip gzipped data: %s", err)
	}

	fmt.Printf("gzip-data: %s", unzippedData)
	return nil
}
