package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type RoomTerrainResponse struct {
	Ok      int               `json:"ok"`
	Terrain []TerrainResponse `json:"terrain"`
}

func (r *RoomTerrainResponse) IsOk() bool {
	return r.Ok == 1
}

type TerrainResponse struct {
	ID      string      `json:"_id"`
	Room    string      `json:"room"`
	X       int         `json:"x"`
	Y       int         `json:"y"`
	Terrain string      `json:"terrain"`
	Type    TerrainType `json:"type"`
}

type TerrainType string

const (
	TerrainTypePlain    TerrainType = "plain"
	TerrainTypeWall                 = "wall"
	TerrainTypeSwamp                = "swamp"
	TerrainTypeAlsoWall             = "also wall"
)

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
