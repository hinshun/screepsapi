package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type XsollaUserResponse struct {
	Ok     int `json:"ok"`
	Active int `json:"active"`
}

func (x *XsollaUserResponse) IsOk() bool {
	return x.Ok == 1
}

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
