package screepsws

import "fmt"

func (ws *WebSocket) SubscribeRoom(shard, room string) (<-chan RoomResponse, error) {
	channel := fmt.Sprintf(roomFormat, shard, room)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan RoomResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := RoomResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *WebSocket) UnsubscribeRoom(shard, room string) error {
	return ws.Unsubscribe(fmt.Sprintf(roomFormat, shard, room))
}

func (ws *WebSocket) SubscribeRoomMap(shard, room string) (<-chan RoomMapResponse, error) {
	channel := fmt.Sprintf(roomMapFormat, shard, room)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan RoomMapResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := RoomMapResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *WebSocket) UnsubscribeRoomMap(shard, room string) error {
	return ws.Unsubscribe(fmt.Sprintf(roomMapFormat, shard, room))
}
