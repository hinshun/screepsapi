package screepsapi

import (
	"fmt"
	"net/http"
)

type WorldStartRoomResponse struct {
	Ok    int      `json:"ok"`
	Rooms []string `json:"room"`
}

func (w *WorldStartRoomResponse) IsOk() bool {
	return w.Ok == 1
}

func (c *Client) WorldStartRoom() (WorldStartRoomResponse, error) {
	worldStartRoomResp := WorldStartRoomResponse{}
	err := c.get(worldStartRoomPath, &worldStartRoomResp, nil, http.StatusOK)
	if err != nil {
		return worldStartRoomResp, fmt.Errorf("failed to get world start room: %s", err)
	}

	return worldStartRoomResp, nil
}
