package screepsws

import (
	"encoding/json"
	"fmt"

	"github.com/hinshun/screepsapi/screepstype"
)

type Response struct {
	Error string
}

type CodeResponse struct {
	Response
	Branch    string            `json:"branch"`
	World     string            `json:"world"`
	Modules   map[string]string `json:"modules"`
	Timestamp int               `json:"timestamp"`
	Hash      int               `json:"hash"`
}

type ConsoleResponse struct {
	Response
	Shard    string                      `json:"shard"`
	Messages screepstype.ConsoleMessages `json:"messages"`
}

type CPUResponse struct {
	Response
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

type MemoryResponse struct {
	Response
}

type MessageResponse struct {
	Response
	NewMessage  screepstype.NewMessage  `json:"newMessage"`
	MessageRead screepstype.MessageRead `json:"messageRead"`
}

type NewMessageResponse struct {
	Response
}

type RoomResponse struct {
	Response
	Flags       int                         `json:"flags"`
	Info        screepstype.RoomInfo        `json:"info"`
	Users       map[string]screepstype.User `json:"users"`
	Controllers map[string]screepstype.Controller
	Creeps      map[string]screepstype.Creep
	Energies    map[string]screepstype.Energy
	Extensions  map[string]screepstype.Extension
	Minerals    map[string]screepstype.Mineral
	Roads       map[string]screepstype.Road
	Sources     map[string]screepstype.Source
	Spawns      map[string]screepstype.Spawn
	Storages    map[string]screepstype.Storage
	Towers      map[string]screepstype.Tower
	Walls       map[string]screepstype.Wall
	Deltas      map[string]json.RawMessage
}

type RoomMapResponse struct {
	Response
	Walls       [][]int `json:"w"`
	Roads       [][]int `json:"r"`
	PowerBanks  [][]int `json:"pb"`
	Portals     [][]int `json:"p"`
	Sources     [][]int `json:"s"`
	Controllers [][]int `json:"c"`
	Minerals    [][]int `json:"m"`
	KeeperLairs [][]int `json:"k"`
	UserObjects map[string][][]int
}

type SetActiveBranchResponse struct {
	Response
	Branch     string                 `json:"branch"`
	ActiveName screepstype.ActiveName `json:"branch"`
}

type ServerMessageResponse struct {
	Response
}

func UnmarshalResponse(data []byte, v interface{}) bool {
	err := json.Unmarshal(data, v)
	if err != nil {
		resp, ok := v.(Response)
		if ok {
			resp.Error = fmt.Sprintf("failed to unmarshal: %s", err)
		}
	}

	return false
}
