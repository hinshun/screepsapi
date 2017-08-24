package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) MapStats(shard string, rooms []string, statName StatName, statsPeriod StatsPeriod) (MapStatsResponse, error) {
	mapStatsReq := MapStatsRequest{
		Shard:    shard,
		Rooms:    rooms,
		StatName: StatName(fmt.Sprintf("%s%s", statName, statsPeriod)),
	}
	mapStatsResp := MapStatsResponse{}

	values := make(url.Values)

	err := c.post(mapStatsPath, &mapStatsReq, &mapStatsResp, values, http.StatusOK)
	if err != nil {
		return mapStatsResp, fmt.Errorf("failed to get map stats: %s", err)
	}

	return mapStatsResp, nil
}
