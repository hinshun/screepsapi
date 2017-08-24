package screepsapi

import (
	"fmt"
	"net/http"
)

type MessagesUnreadCountResponse struct {
	Ok    int `json:"ok"`
	Count int `json:"count"`
}

func (u *MessagesUnreadCountResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) MessagesUnreadCount() (MessagesUnreadCountResponse, error) {
	messagesUnreadCountResp := MessagesUnreadCountResponse{}

	err := c.get(messagesUnreadCountPath, &messagesUnreadCountResp, nil, http.StatusOK)
	if err != nil {
		return messagesUnreadCountResp, fmt.Errorf("failed to get messages unread count: %s", err)
	}

	return messagesUnreadCountResp, nil
}
