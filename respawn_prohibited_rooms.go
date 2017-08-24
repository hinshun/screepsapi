package screepsapi

import (
	"fmt"
	"net/http"
)

type RespawnProhibitedRoomsResponse struct {
	Ok    int      `json:"ok"`
	Rooms []string `json:"rooms"`
}

func (r *RespawnProhibitedRoomsResponse) IsOk() bool {
	return r.Ok == 1
}

func (c *Client) RespawnProhibitedRooms() (RespawnProhibitedRoomsResponse, error) {
	respawnProhibitedRoomsResp := RespawnProhibitedRoomsResponse{}
	err := c.get(respawnProhibitedRoomsPath, &respawnProhibitedRoomsResp, nil, http.StatusOK)
	if err != nil {
		return respawnProhibitedRoomsResp, fmt.Errorf("failed to get user respawn prohibited rooms: %s", err)
	}

	return respawnProhibitedRoomsResp, nil
}
