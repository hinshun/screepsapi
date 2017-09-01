package screepsws

import (
	"encoding/json"
	"fmt"

	"github.com/hinshun/screepsapi/screepstype"
)

type RawRoomResponse struct {
	Response
	Flags   int                         `json:"flags"`
	Info    screepstype.RoomInfo        `json:"info"`
	Users   map[string]screepstype.User `json:"users"`
	Objects map[string]json.RawMessage  `json:"objects"`
}

func (r *RoomResponse) UnmarshalJSON(b []byte) error {
	rawRoomResp := RawRoomResponse{}
	err := json.Unmarshal(b, &rawRoomResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal raw response: %s", err)
	}

	r.Response = rawRoomResp.Response
	r.Flags = rawRoomResp.Flags
	r.Info = rawRoomResp.Info
	r.Users = rawRoomResp.Users
	r.Controllers = make(map[string]screepstype.Controller)
	r.Creeps = make(map[string]screepstype.Creep)
	r.Energies = make(map[string]screepstype.Energy)
	r.Extensions = make(map[string]screepstype.Extension)
	r.Minerals = make(map[string]screepstype.Mineral)
	r.Roads = make(map[string]screepstype.Road)
	r.Sources = make(map[string]screepstype.Source)
	r.Spawns = make(map[string]screepstype.Spawn)
	r.Storages = make(map[string]screepstype.Storage)
	r.Towers = make(map[string]screepstype.Tower)
	r.Walls = make(map[string]screepstype.Wall)
	r.Deltas = make(map[string]json.RawMessage)

	for id, rawObject := range rawRoomResp.Objects {
		objectJSON, err := rawObject.MarshalJSON()
		if err != nil {
			return fmt.Errorf("failed to marshal raw object: %s", err)
		}

		object := screepstype.Object{}
		err = json.Unmarshal(objectJSON, &object)
		if err != nil {
			return fmt.Errorf("failed to unmarshal into object: %s", err)
		}

		switch object.Type {
		case screepstype.ObjectTypeController:
			controller := screepstype.Controller{}
			err = json.Unmarshal(objectJSON, &controller)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into controller: %s", err)
			}
			r.Controllers[id] = controller

		case screepstype.ObjectTypeCreep:
			creep := screepstype.Creep{}
			err = json.Unmarshal(objectJSON, &creep)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into creep: %s", err)
			}
			r.Creeps[id] = creep

		case screepstype.ObjectTypeEnergy:
			energy := screepstype.Energy{}
			err = json.Unmarshal(objectJSON, &energy)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into energy: %s", err)
			}
			r.Energies[id] = energy

		case screepstype.ObjectTypeExtension:
			extension := screepstype.Extension{}
			err = json.Unmarshal(objectJSON, &extension)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into extension: %s", err)
			}
			r.Extensions[id] = extension

		case screepstype.ObjectTypeMineral:
			mineral := screepstype.Mineral{}
			err = json.Unmarshal(objectJSON, &mineral)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into mineral: %s", err)
			}
			r.Minerals[id] = mineral

		case screepstype.ObjectTypeRoad:
			road := screepstype.Road{}
			err = json.Unmarshal(objectJSON, &road)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into road: %s", err)
			}
			r.Roads[id] = road

		case screepstype.ObjectTypeSource:
			source := screepstype.Source{}
			err = json.Unmarshal(objectJSON, &source)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into source: %s", err)
			}
			r.Sources[id] = source

		case screepstype.ObjectTypeSpawn:
			spawn := screepstype.Spawn{}
			err = json.Unmarshal(objectJSON, &spawn)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into spawn: %s", err)
			}
			r.Spawns[id] = spawn

		case screepstype.ObjectTypeStorage:
			storage := screepstype.Storage{}
			err = json.Unmarshal(objectJSON, &storage)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into storage: %s", err)
			}
			r.Storages[id] = storage

		case screepstype.ObjectTypeTower:
			tower := screepstype.Tower{}
			err = json.Unmarshal(objectJSON, &tower)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into tower: %s", err)
			}
			r.Towers[id] = tower

		case screepstype.ObjectTypeWall:
			wall := screepstype.Wall{}
			err = json.Unmarshal(objectJSON, &wall)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into wall: %s", err)
			}
			r.Walls[id] = wall

		default:
			var rawMessage json.RawMessage
			err = json.Unmarshal(objectJSON, &rawMessage)
			if err != nil {
				return fmt.Errorf("failed to unmarshal into raw message: %s", err)
			}
			r.Deltas[id] = rawMessage
		}
	}

	return nil
}
