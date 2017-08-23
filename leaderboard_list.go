package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type LeaderboardListResponse struct {
	Ok    int                                `json:"ok"`
	Ranks []LeaderboardRankResponse          `json:"list"`
	Count int                                `json:"count"`
	Users map[string]LeaderboardUserResponse `json:"users"`
}

type LeaderboardUserResponse struct {
	ID       string        `json:"_id"`
	Username string        `json:"username"`
	Badge    BadgeResponse `json:"badge"`
	GCL      int           `json:"gcl"`
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

	err := c.Get(leaderboardListPath, &leaderboardListResp, values, http.StatusOK)
	if err != nil {
		return leaderboardListResp, fmt.Errorf("failed to get leaderboard list: %s", err)
	}

	return leaderboardListResp, nil
}
