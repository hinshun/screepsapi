package screepsapi

import (
	"fmt"
	"net/http"
)

type MarketMyOrdersResponse struct {
	Ok     int                   `json:"ok"`
	Orders []MarketOrderResponse `json:"list"`
}

func (m *MarketMyOrdersResponse) IsOk() bool {
	return m.Ok == 1
}

func (c *Client) MarketMyOrders() (MarketMyOrdersResponse, error) {
	marketMyOrdersResp := MarketMyOrdersResponse{}
	err := c.get(marketMyOrdersPath, &marketMyOrdersResp, nil, http.StatusOK)
	if err != nil {
		return marketMyOrdersResp, fmt.Errorf("failed to get market orders index: %s", err)
	}

	return marketMyOrdersResp, nil
}
