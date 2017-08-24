package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type MessagesListResponse struct {
	Ok       int               `json:"ok"`
	Messages []MessageResponse `json:"messages"`
}

type MessageResponse struct {
	ID     string `json:"_id"`
	Date   string `json:"date"`
	Type   string `json:"type"`
	Text   string `json:"text"`
	Unread bool   `json:"unread"`
}

func (u *MessagesListResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) MessagesList(respondent string) (MessagesListResponse, error) {
	messagesListResp := MessagesListResponse{}

	values := make(url.Values)
	values.Add(respondentKey, respondent)

	err := c.get(messagesListPath, &messagesListResp, values, http.StatusOK)
	if err != nil {
		return messagesListResp, fmt.Errorf("failed to get messages list: %s", err)
	}

	return messagesListResp, nil
}
