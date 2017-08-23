package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserMessagesIndexResponse struct {
	Ok       int                           `json:"ok"`
	Messages []DetailedUserMessageResponse `json:"messages"`
	Users    map[string]UserResponse       `json:"users"`
}

type DetailedUserMessageResponse struct {
	UserMessageResponse
	User       string `json:"user"`
	Respondent string `json:"respondent"`
}

func (u *UserMessagesIndexResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserMessagesIndex() (UserMessagesIndexResponse, error) {
	userMessagesIndexResp := UserMessagesIndexResponse{}
	err := c.get(userMessagesIndexPath, &userMessagesIndexResp, make(url.Values), http.StatusOK)
	if err != nil {
		return userMessagesIndexResp, fmt.Errorf("failed to get user messages index: %s", err)
	}

	return userMessagesIndexResp, nil
}
