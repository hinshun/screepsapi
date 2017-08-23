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
	CreepsLost1Hour           int `json:creepsLost8"`
	CreepsLost24Hours         int `json:creepsLost180"`
	CreepsLost7Days           int `json:creepsLost1440"`
	CreepsProduced1Hour       int `json:creepsProduced8"`
	CreepsProduced24Hours     int `json:creepsProduced180"`
	CreepsProduced7Days       int `json:creepsProduced1440"`
	Energy1Hour               int `json:energy8"`
	Energy24Hours             int `json:energy180"`
	Energy7Days               int `json:energy1440"`
	EnergyConstruction1Hour   int `json:energyConstruction8"`
	EnergyConstruction24Hours int `json:energyConstruction180"`
	EnergyConstruction7Days   int `json:energyConstruction1440"`
	EnergyControl1Hour        int `json:energyControl8"`
	EnergyControl24Hours      int `json:energyControl180"`
	EnergyControl7Days        int `json:energyControl1440"`
	EnergyCreeps1Hour         int `json:energyCreeps8"`
	EnergyCreeps24Hours       int `json:energyCreeps180"`
	EnergyCreeps7Days         int `json:energyCreeps1440"`
	EnergyHarvested1Hour      int `json:energyHarvested8"`
	EnergyHarvested24Hours    int `json:energyHarvested180"`
	EnergyHarvested7Days      int `json:energyHarvested1440"`
	Power1Hour                int `json:power8"`
	Power24Hours              int `json:power180"`
	Power7Days                int `json:power1440"`
	PowerProcessed1Hour       int `json:powerProcessed8"`
	PowerProcessed24Hours     int `json:powerProcessed180"`
	PowerProcessed7Days       int `json:powerProcessed1440"`
}

func (c *Client) RoomOverview(shard, room string, statsPeriod StatsPeriod) (RoomOverviewResponse, error) {
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
