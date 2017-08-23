package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type MarketOrdersIndexResponse struct {
	Ok     int                   `json:"ok"`
	Orders []MarketOrderResponse `json:"list"`
}

type MarketOrderResponse struct {
	ID    string `json:"_id"`
	Count int    `json:"count"`
}

func (m *MarketOrdersIndexResponse) IsOk() bool {
	return m.Ok == 1
}

func (c *Client) MarketOrdersIndex(shard string) (MarketOrdersIndexResponse, error) {
	marketOrdersIndexResp := MarketOrdersIndexResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.get(marketOrdersIndexPath, &marketOrdersIndexResp, values, http.StatusOK)
	if err != nil {
		return marketOrdersIndexResp, fmt.Errorf("failed to get market orders index: %s", err)
	}

	return marketOrdersIndexResp, nil
}
