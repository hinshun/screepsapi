package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type LeaderboardFindResponse struct {
	RankResponse

	Ok    int            `json:"ok"`
	Ranks []RankResponse `json:"list"`
}

type RankResponse struct {
	ID     string `json:"_id"`
	Season string `json:"season"`
	User   string `json:"user"`
	Score  int    `json:"score"`
	Rank   int    `json:"rank"`
}

func (l *LeaderboardFindResponse) IsOk() bool {
	return l.Ok == 1
}

func (c *Client) LeaderboardFind(mode, username, season string) (LeaderboardFindResponse, error) {
	leaderboardFindResp := LeaderboardFindResponse{}

	values := make(url.Values)
	values.Add(modeKey, mode)
	values.Add(usernameKey, username)
	if season != "" {
		values.Add(seasonKey, season)
	}

	err := c.get(leaderboardFindPath, &leaderboardFindResp, values, http.StatusOK)
	if err != nil {
		return leaderboardFindResp, fmt.Errorf("failed to get leaderboard find: %s", err)
	}

	return leaderboardFindResp, nil
}
