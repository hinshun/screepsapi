package screepsapi

import (
	"fmt"
	"net/http"
)

type UserRoomsResponse struct {
	Ok     int                 `json:"ok"`
	Shards map[string][]string `json:"shards"`
}

func (u *UserRoomsResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserRooms() (UserRoomsResponse, error) {
	userRoomsResp := UserRoomsResponse{}
	err := c.get(userRoomsPath, &userRoomsResp, nil, http.StatusOK)
	if err != nil {
		return userRoomsResp, fmt.Errorf("failed to get user respawn prohibited rooms: %s", err)
	}

	return userRoomsResp, nil
}
