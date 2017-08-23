package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type LeaderboardListResponse struct {
	Ok    int                       `json:"ok"`
	Ranks []LeaderboardRankResponse `json:"list"`
	Count int                       `json:"count"`
	Users map[string]UserResponse   `json:"users"`
}

func (l *LeaderboardListResponse) IsOk() bool {
	return l.Ok == 1
}

func (c *Client) LeaderboardList(mode, season string, limit, offset int) (LeaderboardListResponse, error) {
	leaderboardListResp := LeaderboardListResponse{}

	values := make(url.Values)
	values.Add(modeKey, mode)
	if season != "" {
		values.Add(seasonKey, season)
	}
	values.Add(limitKey, strconv.Itoa(limit))
	values.Add(offsetKey, strconv.Itoa(offset))

	err := c.get(leaderboardListPath, &leaderboardListResp, values, http.StatusOK)
	if err != nil {
		return leaderboardListResp, fmt.Errorf("failed to get leaderboard list: %s", err)
	}

	return leaderboardListResp, nil
}
