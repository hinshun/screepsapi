package screepsapi

import (
	"fmt"
	"net/http"
)

type UserMessagesUnreadCountResponse struct {
	Ok    int `json:"ok"`
	Count int `json:"count"`
}

func (u *UserMessagesUnreadCountResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserMessagesUnreadCount() (UserMessagesUnreadCountResponse, error) {
	userMessagesUnreadCountResp := UserMessagesUnreadCountResponse{}

	err := c.Get(userMessagesUnreadCountPath, &userMessagesUnreadCountResp, nil, http.StatusOK)
	if err != nil {
		return userMessagesUnreadCountResp, fmt.Errorf("failed to get user messages unread count: %s", err)
	}

	return userMessagesUnreadCountResp, nil
}
