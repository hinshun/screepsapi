package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type MarketOrdersResponse struct {
	Ok     int                   `json:"ok"`
	Orders []MarketOrderResponse `json:"list"`
}

func (m *MarketOrdersResponse) IsOk() bool {
	return m.Ok == 1
}

func (c *Client) MarketOrders(shard, resourceType string) (MarketOrdersResponse, error) {
	marketOrdersResp := MarketOrdersResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(resourceTypeKey, resourceType)

	err := c.get(marketOrdersPath, &marketOrdersResp, values, http.StatusOK)
	if err != nil {
		return marketOrdersResp, fmt.Errorf("failed to get market orders: %s", err)
	}

	return marketOrdersResp, nil
}
