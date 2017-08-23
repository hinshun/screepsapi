package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserMessagesIndexResponse struct {
	Ok       int                                     `json:"ok"`
	Messages []UserMessageIndexMessageResponse       `json:"messages"`
	Users    map[string]UserMessageIndexUserResponse `json:"users"`
}

type UserMessageIndexMessageResponse struct {
	UserMessageListMessageResponse
	User       string `json:"user"`
	Respondent string `json:"respondent"`
}

type UserMessageIndexUserResponse struct {
	ID       string        `json:"_id"`
	Username string        `json:"username"`
	Badge    BadgeResponse `json:"badge"`
}

func (u *UserMessagesIndexResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserMessagesIndex() (UserMessagesIndexResponse, error) {
	userMessagesIndexResp := UserMessagesIndexResponse{}
	err := c.Get(userMessagesIndexPath, &userMessagesIndexResp, make(url.Values), http.StatusOK)
	if err != nil {
		return userMessagesIndexResp, fmt.Errorf("failed to get user messages index: %s", err)
	}

	return userMessagesIndexResp, nil
}
