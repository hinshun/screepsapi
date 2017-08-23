package screepsapi

import (
	"fmt"
	"net/http"
)

type UserMessagesSendRequest struct {
	Respondent string `json:"respondent"`
	Text       string `json:"text"`
}

type UserMessagesSendResponse struct {
	Ok int `json:"ok"`
}

func (u *UserMessagesSendResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserMessagesSend(respondent, text string) (UserMessagesSendResponse, error) {
	userMessagesSendReq := UserMessagesSendRequest{
		Respondent: respondent,
		Text:       text,
	}
	userMessagesSendResp := UserMessagesSendResponse{}

	err := c.post(userMessagesSendPath, &userMessagesSendReq, &userMessagesSendResp, nil, http.StatusOK)
	if err != nil {
		return userMessagesSendResp, fmt.Errorf("failed to post user messages send: %s", err)
	}

	return userMessagesSendResp, nil
}
