package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type RoomOverviewResponse struct {
	Ok    int           `json:"ok"`
	Owner OwnerResponse `json:"owner"`
	Stats StatsResponse `json:"stats"`
}

func (r *RoomOverviewResponse) IsOk() bool {
	return r.Ok == 1
}

type OwnerResponse struct {
	Username string        `json:"username"`
	Badge    BadgeResponse `json:"badge"`
}

type StatsResponse struct {
	EnergyHarvested    []StatsPointResponse `json:"energyHarvested"`
	EnergyConstruction []StatsPointResponse `json:"energyConstruction"`
	EnergyCreeps       []StatsPointResponse `json:"energyCreeps"`
	CreepsLost         []StatsPointResponse `json:"creepsLost "`
	StatsMax           StatsMaxResponse     `json:"statsMax"`
	Totals             TotalsResponse       `json:"totals"`
}

type StatsPointResponse struct {
	Value   int `json:"value"`
	EndTime int `json:"endTime"`
}

type StatsMaxResponse struct {
	CreepsLost8            int `json:creepsLost8"`
	CreepsLost180          int `json:creepsLost180"`
	CreepsLost1440         int `json:creepsLost1440"`
	CreepsProduced8        int `json:creepsProduced8"`
	CreepsProduced180      int `json:creepsProduced180"`
	CreepsProduced1440     int `json:creepsProduced1440"`
	Energy8                int `json:energy8"`
	Energy180              int `json:energy180"`
	Energy1440             int `json:energy1440"`
	EnergyConstruction8    int `json:energyConstruction8"`
	EnergyConstruction180  int `json:energyConstruction180"`
	EnergyConstruction1440 int `json:energyConstruction1440"`
	EnergyControl8         int `json:energyControl8"`
	EnergyControl180       int `json:energyControl180"`
	EnergyControl1440      int `json:energyControl1440"`
	EnergyCreeps8          int `json:energyCreeps8"`
	EnergyCreeps180        int `json:energyCreeps180"`
	EnergyCreeps1440       int `json:energyCreeps1440"`
	EnergyHarvested8       int `json:energyHarvested8"`
	EnergyHarvested180     int `json:energyHarvested180"`
	EnergyHarvested1440    int `json:energyHarvested1440"`
	Power8                 int `json:power8"`
	Power180               int `json:power180"`
	Power1440              int `json:power1440"`
	PowerProcessed8        int `json:powerProcessed8"`
	PowerProcessed180      int `json:powerProcessed180"`
	PowerProcessed1440     int `json:powerProcessed1440"`
}

type TotalsResponse struct {
	CreepsProduced     int `json:"creepsProduced"`
	EnergyConstruction int `json:"energyConstruction"`
	EnergyControl      int `json:"energyControl"`
	EnergyCreeps       int `json:"energyCreeps"`
	EnergyHarvested    int `json:"energyHarvested"`
}

func (c *Client) RoomOverview(shard, room, interval string) (RoomOverviewResponse, error) {
	roomOverviewResp := RoomOverviewResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(roomKey, room)
	values.Add(intervalKey, interval)

	err := c.Get(roomOverviewPath, &roomOverviewResp, values, http.StatusOK)
	if err != nil {
		return roomOverviewResp, fmt.Errorf("failed to get room overview: %s", err)
	}

	return roomOverviewResp, nil
}
