package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type ConsoleRequest struct {
	Expression string `json:"expression"`
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
		return insertResp, fmt.Errorf("failed to send console: %s", err)
	}

	return insertResp, nil
}
