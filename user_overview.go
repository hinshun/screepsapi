package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserOverviewResponse struct {
	Ok     int                      `json:"ok"`
	Shards map[string]ShardResponse `json:"shards"`
	Totals TotalsResponse           `json:"totals"`
}

type ShardResponse struct {
	Rooms     []RoomResponse                `json:"rooms"`
	Stats     map[string]StatsPointResponse `json:"stats"`
	GameTimes []*int                        `json:"gametimes"`
}

func (u *UserOverviewResponse) IsOk() bool {
	return u.Ok == 1
}

func (c *Client) UserOverview(statName StatName, statsPeriod StatsPeriod) (UserOverviewResponse, error) {
	userOverviewResp := UserOverviewResponse{}

	values := make(url.Values)
	if statsPeriod != StatsPeriodNone {
		values.Add(intervalKey, string(statsPeriod))
	}
	if statName != StatNameNone {
		values.Add(statNameKey, string(statName))
	}

	err := c.get(userOverviewPath, &userOverviewResp, values, http.StatusOK)
	if err != nil {
		return userOverviewResp, fmt.Errorf("failed to get user find: %s", err)
	}

	return userOverviewResp, nil
}
