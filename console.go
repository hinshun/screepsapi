package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type ConsoleRequest struct {
	Expression string `json:"expression"`
}

type InsertResponse struct {
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

func (c *InsertResponse) IsOk() bool {
	return c.Ok == 1
}

func (c *Client) Console(shard, expression string) (InsertResponse, error) {
	consoleReq := ConsoleRequest{
		Expression: expression,
	}
	insertResp := InsertResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.post(consolePath, &consoleReq, &insertResp, values, http.StatusOK)
	if err != nil {
		return insertResp, fmt.Errorf("failed to post user console: %s", err)
	}

	return insertResp, nil
}
