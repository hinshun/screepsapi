package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type TimeResponse struct {
	Ok int `json:"ok"`
}

func (t *TimeResponse) IsOk() bool {
	return t.Ok == 1
}

func (c *Client) Time(shard string) (TimeResponse, error) {
	timeResp := TimeResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.get(timePath, &timeResp, values, http.StatusOK)
	if err != nil {
		return timeResp, fmt.Errorf("failed to get user stats: %s", err)
	}

	return timeResp, nil
}
