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

type Color int

const (
	ColorNone Color = iota
	ColorRed
	ColorPurple
	ColorBlue
	ColorCyan
	ColorGreen
	ColorYellow
	ColorOrange
	ColorBrown
	ColorGrey
	ColorWhite
)

type ConnectionResponse struct {
	ID   int    `json:"id"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DocumentResponse struct {
	Ok          int `json:"ok"`
	N           int `json:"n"`
	NumModified int `json:"nModified"`
}

type InsertResponse struct {
	Ok            int                  `json:"ok"`
	Result        InsertResultResponse `json:"result"`
	Ops           []InsertOpsResponse  `json:"ops"`
	InsertedCount int                  `json:"insertedCount"`
	InsertedIDs   []string             `json:"insertedIDs"`
}

func (c *InsertResponse) IsOk() bool {
	return c.Ok == 1
}

type InsertResultResponse struct {
	Ok int `json:"ok"`
	N  int `json:"n"`
}

type InsertOpsResponse struct {
	ID         string `json:"_id"`
	User       string `json:"user"`
	Expression string `json:"expression"`
	Hidden     bool   `json:"hidden"`
}

type MarketOrderResponse struct {
	ID              string    `json:"_id"`
	RoomName        string    `json:"roomName"`
	Type            OrderType `json:"type"`
	Price           float32   `json:"price"`
	Amount          int       `json:"amount"`
	RemainingAmount int       `json:"remainingAmount"`
}

type MetaMessageResponse struct {
	Parsed           bool                       `json:"parsed"`
	Index            int                        `json:"index"`
	HashedName       string                     `json:"hashedName"`
	Length           int                        `json:"length"`
	RequestID        int                        `json:"requestId"`
	ResponseTo       int                        `json:"responseTo"`
	ResponseFlags    int                        `json:"responseFlags"`
	CursorID         string                     `json:"cursorId"`
	CursorNotFound   bool                       `json:"cursorNotFound"`
	QueryFailure     bool                       `json:"queryFailure"`
	ShardConfigStale bool                       `json:"shardConfigStale"`
	AwaitCapable     bool                       `json:"awaitCapable"`
	PromoteLongs     bool                       `json:"promoteLongs"`
	PromoteValues    bool                       `json:"promoteValues"`
	PromoteBuffers   bool                       `json:"promoteBuffers"`
	StartingFrom     int                        `json:"startingFrom"`
	NumberReturned   int                        `json:"numberReturned"`
	Raw              MetaMessageDataResponse    `json:"raw"`
	Data             MetaMessageDataResponse    `json:"data"`
	Options          MetaMessageOptionsResponse `json:"opts"`
	Documents        []DocumentResponse         `json:"documents"`
}

type MetaMessageDataResponse struct {
	Type MetaMessageDataType `json:"type"`
	Data []int               `json:"data"`
}

type MetaMessageDataType string

const (
	MetaMessageDataTypeBuffer MetaMessageDataType = "Buffer"
)

type MetaMessageOptionsResponse struct {
	PromoteLongs   bool `json:"promoteLongs"`
	PromoteValues  bool `json:"promoteValues"`
	PromoteBuffers bool `json:"promoteBuffers"`
}

type OrderType string

const (
	OrderTypeBuy  OrderType = "buy"
	OrderTypeSell           = "sell"
)

type StatName string

const (
	StatNameNone               StatName = ""
	StatNameOwner                       = "owner0"
	StatNameClaim                       = "claim0"
	StatNameCreepsLost                  = "creepsLost"
	StatNameCreepsProduced              = "creepsProduced"
	StatNameEnergyConstruction          = "energyConstruction"
	StatNameEnergyControl               = "energyControl"
	StatNameEnergyCreeps                = "energyCreeps"
	StatNameEnergyHarvested             = "energyHarvested"
)

type StatsPeriod string

const (
	StatsPeriodNone    StatsPeriod = ""
	StatsPeriod1Hour               = "8"
	StatsPeriod24Hours             = "180"
	StatsPeriod7Days               = "1440"
)

type TotalsResponse struct {
	PowerProcessed     int `json:"powerProcessed"`
	CreepsLost         int `json:"creepsLost"`
	CreepsProduced     int `json:"creepsProduced"`
	EnergyConstruction int `json:"energyConstruction"`
	EnergyControl      int `json:"energyControl"`
	EnergyCreeps       int `json:"energyCreeps"`
	EnergyHarvested    int `json:"energyHarvested"`
}

type UserResponse struct {
	ID       string        `json:"_id"`
	Username string        `json:"username"`
	Badge    BadgeResponse `json:"badge"`
	GCL      int           `json:"gcl"`
}

type UpsertResponse struct {
	Ok         int                 `json:"ok"`
	Result     DocumentResponse    `json:"result"`
	Connection ConnectionResponse  `json:"connection"`
	Message    MetaMessageResponse `json:"message"`
}

func (c *UpsertResponse) IsOk() bool {
	return c.Ok == 1
}
