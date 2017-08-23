package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserMessagesListResponse struct {
	Ok       int                              `json:"ok"`
	Messages []UserMessageListMessageResponse `json:"messages"`
}

type UserMessageListMessageResponse struct {
	ID     string `json:"_id"`
	Date   string `json:"date"`
	Type   string `json:"type"`
	Text   string `json:"text"`
	Unread bool   `json:"unread"`
}

func (u *UserMessagesListResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserMessagesList(respondent string) (UserMessagesListResponse, error) {
	userMessagesListResp := UserMessagesListResponse{}

	values := make(url.Values)
	values.Add(respondentKey, respondent)

	err := c.Get(userMessagesListPath, &userMessagesListResp, values, http.StatusOK)
	if err != nil {
		return userMessagesListResp, fmt.Errorf("failed to get user messages list: %s", err)
	}

	return userMessagesListResp, nil
}
