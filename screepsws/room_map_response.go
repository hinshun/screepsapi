package screepsws

import (
	"encoding/json"
	"fmt"
)

func (r *RoomMapResponse) UnmarshalJSON(b []byte) error {
	m := make(map[string][][]int)

	err := json.Unmarshal(b, &m)
	if err != nil {
		return fmt.Errorf("failed to unmarshal raw response: %s", err)
	}

	r.UserObjects = make(map[string][][]int)
	for field, coords := range m {
		switch field {
		case "w":
			r.Walls = coords
		case "r":
			r.Roads = coords
		case "pb":
			r.PowerBanks = coords
		case "p":
			r.Portals = coords
		case "s":
			r.Sources = coords
		case "c":
			r.Controllers = coords
		case "m":
			r.Minerals = coords
		case "k":
			r.KeeperLairs = coords
		default:
			r.UserObjects[field] = coords
		}
	}

	return nil
}
