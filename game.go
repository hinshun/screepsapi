package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hinshun/screepsapi/screepstype"
)

func (c *Client) AddObjectIntent(addObjectIntentReq AddObjectIntentRequest) (UpsertResponse, error) {
	upsertResp := UpsertResponse{}
	err := c.post(addObjectIntentPath, &addObjectIntentReq, &upsertResp, nil, http.StatusOK)
	if err != nil {
		return upsertResp, fmt.Errorf("failed to add object intent: %s", err)
	}

	return upsertResp, nil
}

func (c *Client) CheckUniqueObjectName(shard, name string, uniqueObjectNameType screepstype.UniqueObjectName) (Response, error) {
	checkUniqueObjectNameReq := CheckUniqueObjectNameRequest{
		Shard: shard,
		Type:  uniqueObjectNameType,
		Name:  name,
	}
	resp := Response{}

	err := c.post(checkUniqueObjectNamePath, &checkUniqueObjectNameReq, &resp, nil, http.StatusOK)
	if err != nil {
		return resp, fmt.Errorf("failed to check unique object name: %s", err)
	}

	return resp, nil
}

func (c *Client) CreateConstruction(createConstructionReq CreateConstructionRequest) (UpsertResponse, error) {
	upsertResp := UpsertResponse{}
	err := c.post(createConstructionPath, &createConstructionReq, &upsertResp, nil, http.StatusOK)
	if err != nil {
		return upsertResp, fmt.Errorf("failed to add object intent: %s", err)
	}

	return upsertResp, nil
}

func (c *Client) CreateFlag(createFlagReq CreateFlagRequest) (UpsertResponse, error) {
	upsertResp := UpsertResponse{}
	err := c.post(createFlagPath, &createFlagReq, &upsertResp, nil, http.StatusOK)
	if err != nil {
		return upsertResp, fmt.Errorf("failed to create flag: %s", err)
	}

	return upsertResp, nil
}

func (c *Client) DestroyStructure(shard, room, user string, ids []string) (UpsertResponse, error) {
	intents := make([]DestroyStructureIntentRequest, len(ids))
	for _, id := range ids {
		intents = append(intents, DestroyStructureIntentRequest{
			ID:       id,
			RoomName: room,
			User:     user,
		})
	}

	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     "room",
		Name:   screepstype.ObjectIntentDestroyStructure,
		Intent: intents,
	})
}

func (c *Client) GenUniqueObjectName(shard string, uniqueObjectNameType screepstype.UniqueObjectName) (GenUniqueObjectNameResponse, error) {
	genUniqueObjectNameReq := GenUniqueObjectNameRequest{
		Shard: shard,
		Type:  uniqueObjectNameType,
	}
	genUniqueObjectNameResp := GenUniqueObjectNameResponse{}

	err := c.post(genUniqueObjectNamePath, &genUniqueObjectNameReq, &genUniqueObjectNameResp, nil, http.StatusOK)
	if err != nil {
		return genUniqueObjectNameResp, fmt.Errorf("failed to generate unique object name: %s", err)
	}

	return genUniqueObjectNameResp, nil
}

func (c *Client) PlaceSpawn(placeSpawnReq PlaceSpawnRequest) (PlaceSpawnResponse, error) {
	placeSpawnResp := PlaceSpawnResponse{}
	err := c.post(placeSpawnPath, &placeSpawnReq, &placeSpawnResp, nil, http.StatusOK)
	if err != nil {
		return placeSpawnResp, fmt.Errorf("failed to add object intent: %s", err)
	}

	return placeSpawnResp, nil
}

func (c *Client) RemoveFlag(shard, room, name string) (UpsertResponse, error) {
	removeFlagReq := RemoveFlagRequest{
		Shard: shard,
		Room:  room,
		Name:  name,
	}
	upsertResp := UpsertResponse{}

	err := c.post(removeFlagPath, &removeFlagReq, &upsertResp, nil, http.StatusOK)
	if err != nil {
		return upsertResp, fmt.Errorf("failed to remove flag: %s", err)
	}

	return upsertResp, nil
}

func (c *Client) RemoveObject(shard, room, id string) (UpsertResponse, error) {
	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     id,
		Name:   screepstype.ObjectIntentRemove,
		Intent: struct{}{},
	})
}

func (c *Client) Respawn() (RespawnResponse, error) {
	respawnResp := RespawnResponse{}
	err := c.post(respawnPath, nil, &respawnResp, nil, http.StatusOK)
	if err != nil {
		return respawnResp, fmt.Errorf("failed to respawn: %s", err)
	}

	return respawnResp, nil
}

func (c *Client) RespawnProhibitedRooms() (RespawnProhibitedRoomsResponse, error) {
	respawnProhibitedRoomsResp := RespawnProhibitedRoomsResponse{}
	err := c.get(respawnProhibitedRoomsPath, &respawnProhibitedRoomsResp, nil, http.StatusOK)
	if err != nil {
		return respawnProhibitedRoomsResp, fmt.Errorf("failed to get user respawn prohibited rooms: %s", err)
	}

	return respawnProhibitedRoomsResp, nil
}

func (c *Client) SuicideCreep(shard, room, id string) (UpsertResponse, error) {
	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     id,
		Name:   screepstype.ObjectIntentSuicideCreep,
		Intent: struct{}{},
	})
}

func (c *Client) Time(shard string) (TimeResponse, error) {
	timeResp := TimeResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.get(timePath, &timeResp, values, http.StatusOK)
	if err != nil {
		return timeResp, fmt.Errorf("failed to get game time: %s", err)
	}

	return timeResp, nil
}

func (c *Client) UnclaimController(shard, room, id string) (UpsertResponse, error) {
	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     id,
		Name:   screepstype.ObjectIntentUnclaimController,
		Intent: struct{}{},
	})
}
