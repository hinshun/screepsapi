package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hinshun/screepsapi/screepstype"
)

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

func (c *Client) UserOverview(statName screepstype.StatName, statsPeriod screepstype.StatsPeriod) (UserOverviewResponse, error) {
	userOverviewResp := UserOverviewResponse{}

	values := make(url.Values)
	if statsPeriod != screepstype.StatsPeriodNone {
		values.Add(intervalKey, string(statsPeriod))
	}
	if statName != screepstype.StatNameNone {
		values.Add(statNameKey, string(statName))
	}

	err := c.get(userOverviewPath, &userOverviewResp, values, http.StatusOK)
	if err != nil {
		return userOverviewResp, fmt.Errorf("failed to get user find: %s", err)
	}

	return userOverviewResp, nil
}

func (c *Client) UserRooms() (UserRoomsResponse, error) {
	userRoomsResp := UserRoomsResponse{}
	err := c.get(userRoomsPath, &userRoomsResp, nil, http.StatusOK)
	if err != nil {
		return userRoomsResp, fmt.Errorf("failed to get user respawn prohibited rooms: %s", err)
	}

	return userRoomsResp, nil
}

func (c *Client) UserStats(id string, statsPeriod screepstype.StatsPeriod) (UserStatsResponse, error) {
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
