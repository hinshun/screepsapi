package screepsapi

import (
	"fmt"
	"net/http"
	"time"
)

type LeaderboardSeasonsResponse struct {
	Ok      int              `json:"ok"`
	Seasons []SeasonResponse `json:"seasons"`
}

type SeasonResponse struct {
	ID   string    `json:"_id"`
	Name string    `json:"string"`
	Date time.Time `json:"date"`
}

func (l *LeaderboardSeasonsResponse) IsOk() bool {
	return l.Ok == 1
}

func (c *Client) LeaderboardSeasons() (LeaderboardSeasonsResponse, error) {
	leaderboardSeasonsResp := LeaderboardSeasonsResponse{}
	err := c.get(leaderboardSeasonsPath, &leaderboardSeasonsResp, nil, http.StatusOK)
	if err != nil {
		return leaderboardSeasonsResp, fmt.Errorf("failed to get leaderboard season: %s", err)
	}

	return leaderboardSeasonsResp, nil
}
