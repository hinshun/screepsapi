package screepsapi

import (
	"fmt"
	"net/http"
)

func (c *Client) WorldStartRoom() (WorldStartRoomResponse, error) {
	worldStartRoomResp := WorldStartRoomResponse{}
	err := c.get(worldStartRoomPath, &worldStartRoomResp, nil, http.StatusOK)
	if err != nil {
		return worldStartRoomResp, fmt.Errorf("failed to get world start room: %s", err)
	}

	return worldStartRoomResp, nil
}

func (c *Client) WorldStatus() (WorldStatusResponse, error) {
	worldStatusResp := WorldStatusResponse{}
	err := c.get(worldStatusPath, &worldStatusResp, nil, http.StatusOK)
	if err != nil {
		return worldStatusResp, fmt.Errorf("failed to get world start room: %s", err)
	}

	return worldStatusResp, nil
}
