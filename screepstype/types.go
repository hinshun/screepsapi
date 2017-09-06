package screepstype

import (
	"encoding/json"
	"time"
)

type Badge struct {
	// Type is either a BadgePreset or a BadgePath describing a custom
	// badge vector.
	Type json.RawMessage `json:"type"`

	// Color1 is the first color of the badge. Colors can be either all
	// unmarshalled into ColorHex or Color256.
	Color1 json.RawMessage `json:"color1"`

	// Color2 is the secondary color of the badge. Colors can be either all
	// unmarshalled into ColorHex or Color256.
	Color2 json.RawMessage `json:"color2"`

	// Color3 is the tertiary color of the badge. Colors can be either all
	// unmarshalled into ColorHex or Color256.
	Color3 json.RawMessage `json:"color3"`

	Param int  `json:"param"`
	Flip  bool `json:"flip"`
}

type BadgePreset int

type BadgePath struct {
	Path1 string `json:"path1"`
	Path2 string `json:"path2"`
}

type Branch struct {
	ID               string `json:"_id"`
	Branch           string `json:"branch"`
	ActiveSimulation bool   `json:"activeSim"`
	ActiveWorld      bool   `json:"activeWorld"`
}

type ChangeOrderPrice struct {
	OrderID  string `json:"orderId"`
	OldPrice int    `json:"oldPrice"`
	NewPrice int    `json:"newPrice"`
}

type ColorHex string
type Color256 int

type Connection struct {
	ID   int    `json:"id"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ConsoleMessages struct {
	Log     []string `json:"log"`
	Results []string `json:"results"`
}

type Document struct {
	Ok          int `json:"ok"`
	N           int `json:"n"`
	NumModified int `json:"nModified"`
}

type ExtendedOrder struct {
	ExtendOrder ExtendOrder `json:"extendOrder"`
}

type ExtendOrder struct {
	OrderID   string `json:"orderId"`
	AddAmount int    `json:"addAmount"`
}

type FulfilledOrder struct {
	ResourceType   string `json:"resourceType"`
	RoomName       string `json:"roomName"`
	TargetRoomName string `json:"targetRoomName"`
	Price          int    `json:"price"`
	NPC            bool   `json:"npc"`
	Amount         int    `json:"amount"`
}

type GitHub struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type InsertOps struct {
	ID         string `json:"_id"`
	User       string `json:"user"`
	Expression string `json:"expression"`
	Hidden     bool   `json:"hidden"`
}

type InsertResult struct {
	Ok int `json:"ok"`
	N  int `json:"n"`
}

type LeaderboardPosition struct {
	ID     string `json:"_id"`
	Season string `json:"season"`
	User   string `json:"user"`
	Score  int    `json:"score"`
	Rank   int    `json:"rank"`
}

type MapRoomOwner struct {
	User  string `json:"string"`
	Level int    `json:"level"`
}

type MapRoomStats struct {
	Status                        RoomStatus      `json:"status"`
	Novice                        int             `json:"novice"`
	OpenTime                      string          `json:"openTime"`
	RespawnArea                   *int            `json:"respawnArea"`
	Own                           MapRoomOwner    `json:"own"`
	OwnerStat                     []StatsWithUser `json:"owner0"`
	ClaimStat                     []StatsWithUser `json:"claim0"`
	CreepsLost1HourStat           []StatsWithUser `json:creepsLost8"`
	CreepsLost24HoursStat         []StatsWithUser `json:creepsLost180"`
	CreepsLost7DaysStat           []StatsWithUser `json:creepsLost1440"`
	CreepsProduced1HourStat       []StatsWithUser `json:creepsProduced8"`
	CreepsProduced24HoursStat     []StatsWithUser `json:creepsProduced180"`
	CreepsProduced7DaysStat       []StatsWithUser `json:creepsProduced1440"`
	Energy1HourStat               []StatsWithUser `json:energy8"`
	Energy24HoursStat             []StatsWithUser `json:energy180"`
	Energy7DaysStat               []StatsWithUser `json:energy1440"`
	EnergyConstruction1HourStat   []StatsWithUser `json:energyConstruction8"`
	EnergyConstruction24HoursStat []StatsWithUser `json:energyConstruction180"`
	EnergyConstruction7DaysStat   []StatsWithUser `json:energyConstruction1440"`
	EnergyControl1HourStat        []StatsWithUser `json:energyControl8"`
	EnergyControl24HoursStat      []StatsWithUser `json:energyControl180"`
	EnergyControl7DaysStat        []StatsWithUser `json:energyControl1440"`
	EnergyCreeps1HourStat         []StatsWithUser `json:energyCreeps8"`
	EnergyCreeps24HoursStat       []StatsWithUser `json:energyCreeps180"`
	EnergyCreeps7DaysStat         []StatsWithUser `json:energyCreeps1440"`
	EnergyHarvested1HourStat      []StatsWithUser `json:energyHarvested8"`
	EnergyHarvested24HoursStat    []StatsWithUser `json:energyHarvested180"`
	EnergyHarvested7DaysStat      []StatsWithUser `json:energyHarvested1440"`
	Power1HourStat                []StatsWithUser `json:power8"`
	Power24HoursStat              []StatsWithUser `json:power180"`
	Power7DaysStat                []StatsWithUser `json:power1440"`
	PowerProcessed1HourStat       []StatsWithUser `json:powerProcessed8"`
	PowerProcessed24HoursStat     []StatsWithUser `json:powerProcessed180"`
	PowerProcessed7DaysStat       []StatsWithUser `json:powerProcessed1440"`
}

type MarketOrder struct {
	ID              string          `json:"_id"`
	RoomName        string          `json:"roomName"`
	Type            MarketOrderType `json:"type"`
	Price           float32         `json:"price"`
	Amount          int             `json:"amount"`
	RemainingAmount int             `json:"remainingAmount"`
}

type Message struct {
	ID     string `json:"_id"`
	Date   string `json:"date"`
	Type   string `json:"type"`
	Text   string `json:"text"`
	Unread bool   `json:"unread"`
}

type MessageRead struct {
	Message MessageReadUpdate `json:"message"`
}

type MessageReadUpdate struct {
	ID     string `json:"_id"`
	Unread bool   `json:"unread"`
}

type MessageUpdate struct {
	ID         string           `json:"_id"`
	OutMessage string           `json:"outMessage"`
	Text       string           `json:"text"`
	Direction  MessageDirection `json:"type"`
	Unread     bool             `json:"unread"`
	User       string           `json:"user"`
	Respondent string           `json:"respondent"`
}

type MessageWithRecipients struct {
	Message
	User       string `json:"user"`
	Respondent string `json:"respondent"`
}

type MoneyMessage struct {
	ID   string    `json:"_id"`
	Date time.Time `json:"date"`
	Tick bool      `json:"tick"`
	User string    `json:"user"`

	// Type defines which struct to unmarshal `Market` into.
	// If `Type` is `MoneyOrderTypeNew`, use `NewOrder`
	// If `Type` is `MoneyOrderTypeExtended`, use `ExtendedOrder`
	// If `Type` is `MoneyOrderTypeFulfilled`, use `FulfilledOrder`
	// If `Type` is `MoneyOrderTypePriceChange`, use `PriceChange`
	Type MoneyOrderType `json:"type"`

	Balance int             `json:"balance"`
	Change  int             `json:"change"`
	Market  json.RawMessage `json:"market"`
}

type MoneyOrder struct {
	Type         string `json:"type"`
	ResourceType string `json:"resourceType"`
	Price        int    `json:"price"`
	TotalAmount  int    `json:"totalAmount"`
	RoomName     string `json:"roomName"`
}

type NewMessage struct {
	Message MessageUpdate `json:"message"`
}

type NewOrder struct {
	Order MoneyOrder `json:"order"`
}

type PriceChange struct {
	ChangeOrderPrice ChangeOrderPrice `json:"changeOrderPrice"`
}

type ResourceCount struct {
	ID    string `json:"_id"`
	Count int    `json:"count"`
}

type Room struct {
	ID     string     `json:"_id"`
	Status RoomStatus `json:"status"`
	Novice int        `json:"novice"`
}

type RoomInfo struct {
	Mode string `json:"mode"`
}

type RoomOwner struct {
	Username string `json:"username"`
	Badge    Badge  `json:"badge"`
}

type RoomStats struct {
	EnergyHarvested    []StatWithTime     `json:"energyHarvested"`
	EnergyConstruction []StatWithTime     `json:"energyConstruction"`
	EnergyCreeps       []StatWithTime     `json:"energyCreeps"`
	CreepsLost         []StatWithTime     `json:"creepsLost "`
	StatsMax           StatsWithIntervals `json:"statsMax"`
	Totals             Stats              `json:"totals"`
}

type RoomTerrain struct {
	ID   string `json:"_id"`
	Room string `json:"room"`
	X    int    `json:"x"`
	Y    int    `json:"y"`

	// EncodedTerrain is a string of digits representing the terrain from top left
	// to bottom right. This string is empty if `encoded` is false.
	// 0: plain
	// 1: wall
	// 2: swamp
	// 3: also wall
	EncodedTerrain string `json:"terrain"`

	Type Terrain `json:"type"`
}

type Season struct {
	ID   string    `json:"_id"`
	Name string    `json:"string"`
	Date time.Time `json:"date"`
}

type ServerData struct {
	HistoryChunkSize int      `json:"historyChunkSize"`
	Shards           []string `json:"shards"`
}

type Shard struct {
	Rooms     []Room                  `json:"rooms"`
	Stats     map[string]StatWithTime `json:"stats"`
	GameTimes []*int                  `json:"gametimes"`
}

type Stats struct {
	CreepsLost         int `json:"creepsLost"`
	CreepsProduced     int `json:"creepsProduced"`
	EnergyConstruction int `json:"energyConstruction"`
	EnergyControl      int `json:"energyControl"`
	EnergyCreeps       int `json:"energyCreeps"`
	EnergyHarvested    int `json:"energyHarvested"`
	PowerProcessed     int `json:"powerProcessed"`
}

type StatsWithUser struct {
	User  string `json:"string"`
	Value int    `json:"value"`
}

type StatWithTime struct {
	Value   int `json:"value"`
	EndTime int `json:"endTime"`
}

type StatsWithIntervals struct {
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

type Steam struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Ownership   []int  `json:"ownership"`
}

type UpsertMessage struct {
	Parsed           bool                 `json:"parsed"`
	Index            int                  `json:"index"`
	HashedName       string               `json:"hashedName"`
	Length           int                  `json:"length"`
	RequestID        int                  `json:"requestId"`
	ResponseTo       int                  `json:"responseTo"`
	ResponseFlags    int                  `json:"responseFlags"`
	CursorID         string               `json:"cursorId"`
	CursorNotFound   bool                 `json:"cursorNotFound"`
	QueryFailure     bool                 `json:"queryFailure"`
	ShardConfigStale bool                 `json:"shardConfigStale"`
	AwaitCapable     bool                 `json:"awaitCapable"`
	PromoteLongs     bool                 `json:"promoteLongs"`
	PromoteValues    bool                 `json:"promoteValues"`
	PromoteBuffers   bool                 `json:"promoteBuffers"`
	StartingFrom     int                  `json:"startingFrom"`
	NumberReturned   int                  `json:"numberReturned"`
	Raw              UpsertMessageData    `json:"raw"`
	Data             UpsertMessageData    `json:"data"`
	Options          UpsertMessageOptions `json:"opts"`
	Documents        []Document           `json:"documents"`
}

type UpsertMessageData struct {
	Type UpsertMessageDataType `json:"type"`
	Data []byte                `json:"data"`
}

type UpsertMessageOptions struct {
	PromoteLongs   bool `json:"promoteLongs"`
	PromoteValues  bool `json:"promoteValues"`
	PromoteBuffers bool `json:"promoteBuffers"`
}

type User struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Badge    Badge  `json:"badge"`
	GCL      int    `json:"gcl"`
}
