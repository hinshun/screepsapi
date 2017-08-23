package screepsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserMoneyHistoryResponse struct {
	Ok      int                      `json:"ok"`
	Page    int                      `json:"page"`
	HasMore bool                     `json:"hasMore"`
	Orders  []UserMoneyOrderResponse `json:"messages"`
}

type UserMoneyOrderResponse struct {
	ID      string          `json:"_id"`
	Date    time.Time       `json:"date"`
	Tick    bool            `json:"tick"`
	User    string          `json:"user"`
	Type    string          `json:"type"`
	Balance int             `json:"balance"`
	Change  int             `json:"change"`
	Market  json.RawMessage `json:"market"`
}

type NewOrderResponse struct {
	Order OrderResponse `json:"order"`
}

type OrderResponse struct {
	Type         string `json:"type"`
	ResourceType string `json:"resourceType"`
	Price        int    `json:"price"`
	TotalAmount  int    `json:"totalAmount"`
	RoomName     string `json:"roomName"`
}

type ExtendedOrderResponse struct {
	ExtendOrder ExtendOrderResponse `json:"extendOrder"`
}

type ExtendOrderResponse struct {
	OrderID   string `json:"orderId"`
	AddAmount int    `json:"addAmount"`
}

type FulfilledOrderResponse struct {
	ResourceType   string `json:"resourceType"`
	RoomName       string `json:"roomName"`
	TargetRoomName string `json:"targetRoomName"`
	Price          int    `json:"price"`
	NPC            bool   `json:"npc"`
	Amount         int    `json:"amount"`
}

type PriceChangeResponse struct {
	ChangeOrderPrice ChangeOrderPriceResponse `json:"changeOrderPrice"`
}

type ChangeOrderPriceResponse struct {
	OrderID  string `json:"orderId"`
	OldPrice int    `json:"oldPrice"`
	NewPrice int    `json:"newPrice"`
}

func (u *UserMoneyHistoryResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserMoneyHistory() (UserMoneyHistoryResponse, error) {
	userMoneyHistoryResp := UserMoneyHistoryResponse{}
	err := c.get(userMoneyHistoryPath, &userMoneyHistoryResp, nil, http.StatusOK)
	if err != nil {
		return userMoneyHistoryResp, fmt.Errorf("failed to get user messages list: %s", err)
	}

	return userMoneyHistoryResp, nil
}
