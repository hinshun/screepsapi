package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type MessagesIndexResponse struct {
	Ok       int                       `json:"ok"`
	Messages []DetailedMessageResponse `json:"messages"`
	Users    map[string]UserResponse   `json:"users"`
}

type DetailedMessageResponse struct {
	MessageResponse
	User       string `json:"user"`
	Respondent string `json:"respondent"`
}

func (m *MessagesIndexResponse) IsOk() bool {
	return m.Ok == 1
}

func (c *Client) MessagesIndex() (MessagesIndexResponse, error) {
	messagesIndexResp := MessagesIndexResponse{}
	err := c.get(messagesIndexPath, &messagesIndexResp, make(url.Values), http.StatusOK)
	if err != nil {
		return messagesIndexResp, fmt.Errorf("failed to get messages index: %s", err)
	}

	return messagesIndexResp, nil
}
