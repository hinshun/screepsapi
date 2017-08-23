package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserConsoleRequest struct {
	Expression string `json:"expression"`
}

type UserConsoleResponse struct {
	Ok            int                       `json:"ok"`
	Result        UserConsoleResultResponse `json:"result"`
	Ops           []UserConsoleOpsResponse  `json:"ops"`
	InsertedCount int                       `json:"insertedCount"`
	InsertedIDs   []string                  `json:"insertedIDs"`
}

type UserConsoleResultResponse struct {
	Ok int `json:"ok"`
	N  int `json:"n"`
}

type UserConsoleOpsResponse struct {
	ID         string `json:"_id"`
	User       string `json:"user"`
	Expression string `json:"expression"`
	Hidden     bool   `json:"hidden"`
}

func (u *UserConsoleResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserConsole(shard, expression string) (UserConsoleResponse, error) {
	userConsoleReq := UserConsoleRequest{
		Expression: expression,
	}
	userConsoleResp := UserConsoleResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.post(userConsolePath, &userConsoleReq, &userConsoleResp, values, http.StatusOK)
	if err != nil {
		return userConsoleResp, fmt.Errorf("failed to post user console: %s", err)
	}

	return userConsoleResp, nil
}
