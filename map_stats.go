package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type MapStatsRequest struct {
	Shard    string   `json:"shard"`
	Rooms    []string `json:"rooms"`
	StatName string   `json:"statName"`
}

type MapStatsResponse struct {
	Ok    int                             `json:"ok"`
	Stats map[string]MapStatsRoomResponse `json:"stats"`
	Users map[string]UserResponse         `json:"users"`
}

type MapStatsRoomResponse struct {
	Status                        RoomStatus                 `json:"status"`
	Novice                        int                        `json:"novice"`
	OpenTime                      string                     `json:"openTime"`
	RespawnArea                   *int                       `json:"respawnArea"`
	Own                           MapStatsRoomOwnResponse    `json:"own"`
	OwnerStat                     []MapStatsRoomStatResponse `json:"owner0"`
	ClaimStat                     []MapStatsRoomStatResponse `json:"claim0"`
	CreepsLost1HourStat           []MapStatsRoomStatResponse `json:creepsLost8"`
	CreepsLost24HoursStat         []MapStatsRoomStatResponse `json:creepsLost180"`
	CreepsLost7DaysStat           []MapStatsRoomStatResponse `json:creepsLost1440"`
	CreepsProduced1HourStat       []MapStatsRoomStatResponse `json:creepsProduced8"`
	CreepsProduced24HoursStat     []MapStatsRoomStatResponse `json:creepsProduced180"`
	CreepsProduced7DaysStat       []MapStatsRoomStatResponse `json:creepsProduced1440"`
	Energy1HourStat               []MapStatsRoomStatResponse `json:energy8"`
	Energy24HoursStat             []MapStatsRoomStatResponse `json:energy180"`
	Energy7DaysStat               []MapStatsRoomStatResponse `json:energy1440"`
	EnergyConstruction1HourStat   []MapStatsRoomStatResponse `json:energyConstruction8"`
	EnergyConstruction24HoursStat []MapStatsRoomStatResponse `json:energyConstruction180"`
	EnergyConstruction7DaysStat   []MapStatsRoomStatResponse `json:energyConstruction1440"`
	EnergyControl1HourStat        []MapStatsRoomStatResponse `json:energyControl8"`
	EnergyControl24HoursStat      []MapStatsRoomStatResponse `json:energyControl180"`
	EnergyControl7DaysStat        []MapStatsRoomStatResponse `json:energyControl1440"`
	EnergyCreeps1HourStat         []MapStatsRoomStatResponse `json:energyCreeps8"`
	EnergyCreeps24HoursStat       []MapStatsRoomStatResponse `json:energyCreeps180"`
	EnergyCreeps7DaysStat         []MapStatsRoomStatResponse `json:energyCreeps1440"`
	EnergyHarvested1HourStat      []MapStatsRoomStatResponse `json:energyHarvested8"`
	EnergyHarvested24HoursStat    []MapStatsRoomStatResponse `json:energyHarvested180"`
	EnergyHarvested7DaysStat      []MapStatsRoomStatResponse `json:energyHarvested1440"`
	Power1HourStat                []MapStatsRoomStatResponse `json:power8"`
	Power24HoursStat              []MapStatsRoomStatResponse `json:power180"`
	Power7DaysStat                []MapStatsRoomStatResponse `json:power1440"`
	PowerProcessed1HourStat       []MapStatsRoomStatResponse `json:powerProcessed8"`
	PowerProcessed24HoursStat     []MapStatsRoomStatResponse `json:powerProcessed180"`
	PowerProcessed7DaysStat       []MapStatsRoomStatResponse `json:powerProcessed1440"`
}
type MapStatsRoomOwnResponse struct {
	User  string `json:"string"`
	Level int    `json:"level"`
}

type MapStatsRoomStatResponse struct {
	User  string `json:"string"`
	Value int    `json:"value"`
}

func (g *MapStatsResponse) IsOk() bool {
	return g.Ok == 1
}

func (c *Client) MapStats(shard string, rooms []string, statName StatName, statsPeriod StatsPeriod) (MapStatsResponse, error) {
	mapStatsReq := MapStatsRequest{
		Shard:    shard,
		Rooms:    rooms,
		StatName: fmt.Sprintf("%s%s", statName, statsPeriod),
	}
	mapStatsResp := MapStatsResponse{}

	values := make(url.Values)

	err := c.post(mapStatsPath, &mapStatsReq, &mapStatsResp, values, http.StatusOK)
	if err != nil {
		return mapStatsResp, fmt.Errorf("failed to get map stats: %s", err)
	}

	return mapStatsResp, nil
}
