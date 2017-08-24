package screepsapi

import (
	"fmt"
	"net/http"
)

type WorldStatusResponse struct {
	Ok     int         `json:"ok"`
	Status WorldStatus `json:"status"`
}

type WorldStatus string

const (
	WorldStatusEmpty  WorldStatus = "empty"
	WorldStatusNormal             = "normal"
)

func (w *WorldStatusResponse) IsOk() bool {
	return w.Ok == 1
}

func (c *Client) WorldStatus() (WorldStatusResponse, error) {
	worldStatusResp := WorldStatusResponse{}
	err := c.get(worldStatusPath, &worldStatusResp, nil, http.StatusOK)
	if err != nil {
		return worldStatusResp, fmt.Errorf("failed to get world start room: %s", err)
	}

	return worldStatusResp, nil
}
