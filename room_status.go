package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type RoomStatusResponse struct {
	Ok   int          `json:"ok"`
	Room RoomResponse `json:"room"`
}

type RoomResponse struct {
	ID     string     `json:"_id"`
	Status RoomStatus `json:"status"`
	Novice int        `json:"novice"`
}

type RoomStatus string

const (
	StatusNormal       RoomStatus = "normal"
	StatusOutOfBorders            = "out of borders"
)

func (r *RoomStatusResponse) IsOk() bool {
	return r.Ok == 1
}

func (c *Client) RoomStatus(shard, room string) (RoomStatusResponse, error) {
	roomStatusResp := RoomStatusResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(roomKey, room)

	err := c.get(roomStatusPath, &roomStatusResp, values, http.StatusOK)
	if err != nil {
		return roomStatusResp, fmt.Errorf("failed to get room status: %s", err)
	}

	return roomStatusResp, nil
}
