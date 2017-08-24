package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) MessagesIndex() (MessagesIndexResponse, error) {
	messagesIndexResp := MessagesIndexResponse{}
	err := c.get(messagesIndexPath, &messagesIndexResp, make(url.Values), http.StatusOK)
	if err != nil {
		return messagesIndexResp, fmt.Errorf("failed to get messages index: %s", err)
	}

	return messagesIndexResp, nil
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

func (c *Client) MessagesSend(respondent, text string) (Response, error) {
	messagesSendReq := MessagesSendRequest{
		Respondent: respondent,
		Text:       text,
	}
	resp := Response{}

	err := c.post(messagesSendPath, &messagesSendReq, &resp, nil, http.StatusOK)
	if err != nil {
		return resp, fmt.Errorf("failed to post messages send: %s", err)
	}

	return resp, nil
}

func (c *Client) MessagesUnreadCount() (MessagesUnreadCountResponse, error) {
	messagesUnreadCountResp := MessagesUnreadCountResponse{}

	err := c.get(messagesUnreadCountPath, &messagesUnreadCountResp, nil, http.StatusOK)
	if err != nil {
		return messagesUnreadCountResp, fmt.Errorf("failed to get messages unread count: %s", err)
	}

	return messagesUnreadCountResp, nil
}
