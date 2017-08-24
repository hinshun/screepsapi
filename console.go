package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type ConsoleRequest struct {
	Expression string `json:"expression"`
}

type ConsoleResponse struct {
	Ok            int                   `json:"ok"`
	Result        ConsoleResultResponse `json:"result"`
	Ops           []ConsoleOpsResponse  `json:"ops"`
	InsertedCount int                   `json:"insertedCount"`
	InsertedIDs   []string              `json:"insertedIDs"`
}

type ConsoleResultResponse struct {
	Ok int `json:"ok"`
	N  int `json:"n"`
}

type ConsoleOpsResponse struct {
	ID         string `json:"_id"`
	User       string `json:"user"`
	Expression string `json:"expression"`
	Hidden     bool   `json:"hidden"`
}

func (u *ConsoleResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) Console(shard, expression string) (ConsoleResponse, error) {
	consoleReq := ConsoleRequest{
		Expression: expression,
	}
	consoleResp := ConsoleResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.post(consolePath, &consoleReq, &consoleResp, values, http.StatusOK)
	if err != nil {
		return consoleResp, fmt.Errorf("failed to post user console: %s", err)
	}

	return consoleResp, nil
}
