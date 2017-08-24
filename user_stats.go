package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserStatsResponse struct {
	Ok    int            `json:"ok"`
	Stats TotalsResponse `json:"stats"`
}

func (u *UserStatsResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserStats(id string, statsPeriod StatsPeriod) (UserStatsResponse, error) {
	userStatsResp := UserStatsResponse{}

	values := make(url.Values)
	values.Add(idKey, id)
	values.Add(intervalKey, string(statsPeriod))

	err := c.get(userStatsPath, &userStatsResp, values, http.StatusOK)
	if err != nil {
		return userStatsResp, fmt.Errorf("failed to get user stats: %s", err)
	}

	return userStatsResp, nil
}
