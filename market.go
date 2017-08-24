package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) MoneyHistory() (MoneyHistoryResponse, error) {
	moneyHistoryResp := MoneyHistoryResponse{}
	err := c.get(moneyHistoryPath, &moneyHistoryResp, nil, http.StatusOK)
	if err != nil {
		return moneyHistoryResp, fmt.Errorf("failed to get money history: %s", err)
	}

	return moneyHistoryResp, nil
}

func (c *Client) Orders(shard, resourceType string) (OrdersResponse, error) {
	ordersResp := OrdersResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(resourceTypeKey, resourceType)

	err := c.get(ordersPath, &ordersResp, values, http.StatusOK)
	if err != nil {
		return ordersResp, fmt.Errorf("failed to get market orders: %s", err)
	}

	return ordersResp, nil
}

func (c *Client) OrdersIndex(shard string) (OrdersIndexResponse, error) {
	ordersIndexResp := OrdersIndexResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.get(ordersIndexPath, &ordersIndexResp, values, http.StatusOK)
	if err != nil {
		return ordersIndexResp, fmt.Errorf("failed to get market orders index: %s", err)
	}

	return ordersIndexResp, nil
}

func (c *Client) MyOrders() (OrdersResponse, error) {
	ordersResp := OrdersResponse{}
	err := c.get(myOrdersPath, &ordersResp, nil, http.StatusOK)
	if err != nil {
		return ordersResp, fmt.Errorf("failed to get market orders index: %s", err)
	}

	return ordersResp, nil
}
