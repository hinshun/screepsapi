package screepsapi

import "encoding/json"

type ScreepsAPIResponse interface {
	IsOk() bool
}

type UserResponse struct {
	ID       string        `json:"_id"`
	Username string        `json:"username"`
	Badge    BadgeResponse `json:"badge"`
	GCL      int           `json:"gcl"`
}

type BadgeResponse struct {
	Type   json.RawMessage `json:"type"`
	Color1 json.RawMessage `json:"color1"`
	Color2 json.RawMessage `json:"color2"`
	Color3 json.RawMessage `json:"color3"`
	Param  int             `json:"param"`
	Flip   bool            `json:"flip"`
}

type BadgeTypeBaseResponse int

type BadgeTypePathResponse struct {
	Path1 string `json:"path1"`
	Path2 string `json:"path2"`
}

type BadgeColorHexResponse string
type BadgeColor256Response int

type StatsPeriod string

const (
	StatsPeriodNone    StatsPeriod = ""
	StatsPeriod1Hour               = "8"
	StatsPeriod24Hours             = "180"
	StatsPeriod7Days               = "1440"
)

type TotalsResponse struct {
	CreepsProduced     int `json:"creepsProduced"`
	EnergyConstruction int `json:"energyConstruction"`
	EnergyControl      int `json:"energyControl"`
	EnergyCreeps       int `json:"energyCreeps"`
	EnergyHarvested    int `json:"energyHarvested"`
}
