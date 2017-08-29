package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hinshun/screepsapi/screepstype"
)

func (c *Client) RoomOverview(shard, room string, statsPeriod screepstype.StatsPeriod) (RoomOverviewResponse, error) {
	roomOverviewResp := RoomOverviewResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(roomKey, room)
	values.Add(intervalKey, string(statsPeriod))

	err := c.get(roomOverviewPath, &roomOverviewResp, values, http.StatusOK)
	if err != nil {
		return roomOverviewResp, fmt.Errorf("failed to get room overview: %s", err)
	}

	return roomOverviewResp, nil
}

func (c *Client) RoomStatus(shard, room string) (RoomStatusResponse, error) {
	roomStatusResp := RoomStatusResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(roomKey, room)

	err := c.get(roomStatusPath, &roomStatusResp, values, http.StatusOK)
	if err != nil {
		return roomStatusResp, fmt.Errorf("failed to get room status: %s", err)
	}

	return roomStatusResp, nil
}

func (c *Client) RoomTerrain(shard, room string, encoded bool) (RoomTerrainResponse, error) {
	roomTerrainResp := RoomTerrainResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(roomKey, room)
	if encoded {
		values.Add(encodedKey, queryTrue)
	}

	err := c.get(roomTerrainPath, &roomTerrainResp, values, http.StatusOK)
	if err != nil {
		return roomTerrainResp, fmt.Errorf("failed to get room terrain: %s", err)
	}

	return roomTerrainResp, nil
}
