package screepsapi

import (
	"fmt"
	"net/http"
)

type MessagesSendRequest struct {
	Respondent string `json:"respondent"`
	Text       string `json:"text"`
}

type MessagesSendResponse struct {
	Ok int `json:"ok"`
}

func (m *MessagesSendResponse) IsOk() bool {
	return m.Ok == 1
}

func (c *Client) MessagesSend(respondent, text string) (MessagesSendResponse, error) {
	messagesSendReq := MessagesSendRequest{
		Respondent: respondent,
		Text:       text,
	}
	messagesSendResp := MessagesSendResponse{}

	err := c.post(messagesSendPath, &messagesSendReq, &messagesSendResp, nil, http.StatusOK)
	if err != nil {
		return messagesSendResp, fmt.Errorf("failed to post messages send: %s", err)
	}

	return messagesSendResp, nil
}
