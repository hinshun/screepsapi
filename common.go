package screepsapi

import "encoding/json"

type ScreepsAPIResponse interface {
	IsOk() bool
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
