package screepsapi

import (
	"fmt"
	"net/http"
)

type AddObjectIntentRequest struct {
	Shard  string       `json:"shard"`
	Room   string       `json:"room"`
	ID     string       `json:"_id"`
	Name   ObjectIntent `json:"name"`
	Intent interface{}  `json:"intent"`
}

type ObjectIntent string

const (
	ObjectIntentRemove            ObjectIntent = "remove"
	ObjectIntentSuicideCreep                   = "suicide"
	ObjectIntentUnclaimController              = "unclaim"
	ObjectIntentDestroyStructure               = "destroyStructure"
)

type DestroyStructureIntentRequest struct {
	ID       string `json:"id"`
	RoomName string `json:"roomName"`
	User     string `json:"user"`
}

func (c *Client) RemoveObject(shard, room, id string) (UpsertResponse, error) {
	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     id,
		Name:   ObjectIntentRemove,
		Intent: struct{}{},
	})
}

func (c *Client) SuicideCreep(shard, room, id string) (UpsertResponse, error) {
	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     id,
		Name:   ObjectIntentSuicideCreep,
		Intent: struct{}{},
	})
}

func (c *Client) UnclaimController(shard, room, id string) (UpsertResponse, error) {
	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     id,
		Name:   ObjectIntentUnclaimController,
		Intent: struct{}{},
	})
}

func (c *Client) DestroyStructure(shard, room, user string, ids []string) (UpsertResponse, error) {
	intents := make([]DestroyStructureIntentRequest, len(ids))
	for _, id := range ids {
		intents = append(intents, DestroyStructureIntentRequest{
			ID:   id,
			Room: room,
			User: user,
		})
	}

	return c.AddObjectIntent(AddObjectIntentRequest{
		Shard:  shard,
		Room:   room,
		ID:     "room",
		Name:   ObjectIntentDestroyStructure,
		Intent: intents,
	})
}

func (c *Client) AddObjectIntent(addObjectIntentReq AddObjectIntentRequest) (UpsertResponse, error) {
	upsertResp := UpsertResponse{}
	err := c.post(addObjectIntentPath, &addObjectIntentReq, &upsertResp, nil, http.StatusOK)
	if err != nil {
		return upsertResp, fmt.Errorf("failed to add object intent: %s", err)
	}

	return upsertResp, nil
}
