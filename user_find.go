package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserFindResponse struct {
	Ok   int          `json:"ok"`
	User UserResponse `json:"user"`
}

func (u *UserFindResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserFind(id, username string) (UserFindResponse, error) {
	userFindResp := UserFindResponse{}

	values := make(url.Values)
	if id != "" {
		values.Add(idKey, id)
	}
	if username != "" {
		values.Add(usernameKey, username)
	}

	err := c.get(userFindPath, &userFindResp, values, http.StatusOK)
	if err != nil {
		return userFindResp, fmt.Errorf("failed to get user find: %s", err)
	}

	return userFindResp, nil
}
