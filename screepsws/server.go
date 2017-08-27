package screepsws

import "fmt"

func (ws *WebSocket) SubscribeServerMessage() (<-chan ServerMessageResponse, error) {
	dataChan, err := ws.Subscribe(serverMessageFormat)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", serverMessageFormat)
	}

	respChan := make(chan ServerMessageResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := ServerMessageResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *WebSocket) UnsubscribeServerMessage() error {
	return ws.Unsubscribe(serverMessageFormat)
}
