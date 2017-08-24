package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) XsollaUser(shard string) (XsollaUserResponse, error) {
	xsollaUserResp := XsollaUserResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.get(xsollaUserPath, &xsollaUserResp, values, http.StatusOK)
	if err != nil {
		return xsollaUserResp, fmt.Errorf("failed to get xsolla user: %s", err)
	}

	return xsollaUserResp, nil
}
