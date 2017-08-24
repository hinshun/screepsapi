package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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

func (c *Client) LeaderboardSeasons() (LeaderboardSeasonsResponse, error) {
	leaderboardSeasonsResp := LeaderboardSeasonsResponse{}
	err := c.get(leaderboardSeasonsPath, &leaderboardSeasonsResp, nil, http.StatusOK)
	if err != nil {
		return leaderboardSeasonsResp, fmt.Errorf("failed to get leaderboard season: %s", err)
	}

	return leaderboardSeasonsResp, nil
}
