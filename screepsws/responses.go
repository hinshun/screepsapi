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
	Shard    string                       `json:"shard"`
	Messages []screepstype.ConsoleMessage `json:"messages"`
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
	Objects  map[string]string    `json:"objects"`
	GameTime int                  `json:"gameTime"`
	Info     screepstype.RoomInfo `json:"info"`
	Visual   string               `json:"visual"`
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
