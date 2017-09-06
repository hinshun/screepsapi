package screepsws

import "fmt"

func (ws *webSocket) SubscribeCode(userID string) (<-chan CodeResponse, error) {
	channel := fmt.Sprintf(codeFormat, userID)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan CodeResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := CodeResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeCode(userID string) error {
	return ws.Unsubscribe(fmt.Sprintf(codeFormat, userID))
}

func (ws *webSocket) SubscribeConsole(userID string) (<-chan ConsoleResponse, error) {
	channel := fmt.Sprintf(consoleFormat, userID)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan ConsoleResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := ConsoleResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeConsole(userID string) error {
	return ws.Unsubscribe(fmt.Sprintf(consoleFormat, userID))
}

func (ws *webSocket) SubscribeCPU(userID string) (<-chan CPUResponse, error) {
	channel := fmt.Sprintf(cpuFormat, userID)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan CPUResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := CPUResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeCPU(userID string) error {
	return ws.Unsubscribe(fmt.Sprintf(cpuFormat, userID))
}

func (ws *webSocket) SubscribeMemory(userID, path string) (<-chan MemoryResponse, error) {
	channel := fmt.Sprintf(memoryFormat, userID, path)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan MemoryResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := MemoryResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeMemory(userID, path string) error {
	return ws.Unsubscribe(fmt.Sprintf(memoryFormat, userID, path))
}

func (ws *webSocket) SubscribeMessage(userID, respondentID string) (<-chan MessageResponse, error) {
	channel := fmt.Sprintf(messageFormat, userID, respondentID)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan MessageResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := MessageResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeMessage(userID, respondentID string) error {
	return ws.Unsubscribe(fmt.Sprintf(messageFormat, userID, respondentID))
}

func (ws *webSocket) SubscribeMoney(userID string) (<-chan int, error) {
	channel := fmt.Sprintf(moneyFormat, userID)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan int)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			var resp int
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeMoney(userID string) error {
	return ws.Unsubscribe(fmt.Sprintf(moneyFormat, userID))
}

func (ws *webSocket) SubscribeNewMessage(userID string) (<-chan NewMessageResponse, error) {
	channel := fmt.Sprintf(newMessageFormat, userID)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan NewMessageResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := NewMessageResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeNewMessage(userID string) error {
	return ws.Unsubscribe(fmt.Sprintf(newMessageFormat, userID))
}

func (ws *webSocket) SubscribeSetActiveBranch(userID string) (<-chan SetActiveBranchResponse, error) {
	channel := fmt.Sprintf(setActiveBranchFormat, userID)
	dataChan, err := ws.Subscribe(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to '%s'", channel)
	}

	respChan := make(chan SetActiveBranchResponse)
	go func() {
		defer close(respChan)
		for data := range dataChan {
			resp := SetActiveBranchResponse{}
			closed := UnmarshalResponse(data, &resp)
			if closed {
				return
			}
			respChan <- resp
		}
	}()

	return respChan, nil
}

func (ws *webSocket) UnsubscribeSetActiveBranch(userID string) error {
	return ws.Unsubscribe(fmt.Sprintf(setActiveBranchFormat, userID))
}
