package screepsapi

import (
	"fmt"
	"net/http"
)

type RemoveFlagRequest struct {
	Shard string `json:"shard"`
	Room  string `json:"room"`
	Name  string `json:"name"`
}

func (c *Client) RemoveFlag(shard, room, name string) (UpsertResponse, error) {
	removeFlagReq := RemoveFlagRequest{
		Shard: shard,
		Room:  room,
		Name:  name,
	}
	upsertResp := UpsertResponse{}

	err := c.post(removeFlagPath, &removeFlagReq, &upsertResp, nil, http.StatusOK)
	if err != nil {
		return upsertResp, fmt.Errorf("failed to remove flag: %s", err)
	}

	return upsertResp, nil
}
