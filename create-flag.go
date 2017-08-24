package screepsapi

import (
	"fmt"
	"net/http"
)

type CreateFlagRequest struct {
	Shard          string `json:"shard"`
	Room           string `json:"room"`
	X              int    `json:"x"`
	Y              int    `json:"y"`
	Name           string `json:"name"`
	Color          Color  `json:"color"`
	SecondaryColor Color  `json:"secondaryColor"`
}

func (c *Client) CreateFlag(createFlagReq CreateFlagRequest) (UpsertResponse, error) {
	upsertResp := UpsertResponse{}
	err := c.post(createFlagPath, &createFlagReq, &upsertResp, nil, http.StatusOK)
	if err != nil {
		return upsertResp, fmt.Errorf("failed to create flag: %s", err)
	}

	return upsertResp, nil
}
