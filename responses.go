package screepsapi

import "encoding/json"

type APIResponse interface {
	IsOk() bool
}

type Response struct {
	Ok    int   `json:"ok"`
	Error Error `json:"error"`
}

func (r *Response) IsOk() bool {
	return r.Ok == 1
}

type BranchesResponse struct {
	Response
	Branches []Branch `json:"list"`
}

type BadgeResponse struct {
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

type GenUniqueObjectNameResponse struct {
	Response
	Name string `json:"name"`
}

type InsertResponse struct {
	Response
	Result        InsertResult `json:"result"`
	Ops           []InsertOps  `json:"ops"`
	InsertedCount int          `json:"insertedCount"`
	InsertedIDs   []string     `json:"insertedIDs"`
}

type LeaderboardFindResponse struct {
	Response
	LeaderboardPosition
	Positions []LeaderboardPosition `json:"list"`
}

type LeaderboardListResponse struct {
	Response
	Positions []LeaderboardPosition `json:"list"`
	Count     int                   `json:"count"`
	Users     map[string]User       `json:"users"`
}

type LeaderboardSeasonsResponse struct {
	Response
	Seasons []Season `json:"seasons"`
}

type MapStatsResponse struct {
	Response
	Stats map[string]MapRoomStats `json:"stats"`
	Users map[string]User         `json:"users"`
}

type MeResponse struct {
	Response
	ID                 string        `json:"_id"`
	Email              string        `json:"email"`
	Username           string        `json:"username"`
	CPU                int           `json:"cpu"`
	Badge              BadgeResponse `json:"badge"`
	Password           bool          `json:"password"`
	Credits            int           `json:"credits"`
	PromoPeriodUntil   int           `json:"promoPeriodUntil"`
	Money              int           `json:"money"`
	SubscriptionTokens int           `json:"subscriptionTokens"`
	GitHub             GitHub        `json:"github"`
	Steam              Steam         `json:"steam"`
}

type MemoryResponse struct {
	Response
	Data interface{} `json:"data"`
}

type MemorySegmentResponse struct {
	Response
	Data string `json:"data"`
}

type MessagesIndexResponse struct {
	Response
	Messages []MessageWithRecipients `json:"messages"`
	Users    map[string]User         `json:"users"`
}

type MessagesListResponse struct {
	Response
	Messages []Message `json:"messages"`
}

type MessagesUnreadCountResponse struct {
	Response
	Count int `json:"count"`
}

type MoneyHistoryResponse struct {
	Response
	Page     int            `json:"page"`
	HasMore  bool           `json:"hasMore"`
	Messages []MoneyMessage `json:"messages"`
}

type MyOrdersResponse struct {
	Response
	Orders []MarketOrder `json:"list"`
}

type OrdersIndexResponse struct {
	Response
	ResourceCounts []ResourceCount `json:"list"`
}

type OrdersResponse struct {
	Response
	Orders []MarketOrder `json:"list"`
}

type PlaceSpawnResponse struct {
	Response
	Newbie bool `json:"newbie"`
}

type PullCodeResponse struct {
	Response
	Branch  string            `json:"branch"`
	Modules map[string]string `json:"modules"`
}

type PushCodeResponse struct {
	Response
	Timestamp int `json:"timestamp"`
	Hash      int `json:"_hash"`
}

type RespawnResponse struct {
	Response
	First  Document `json:"0"`
	Second Document `json:"1"`
}

type RespawnProhibitedRoomsResponse struct {
	Response
	Rooms []string `json:"rooms"`
}

type RoomOverviewResponse struct {
	Response
	Owner RoomOwner `json:"owner"`
	Stats RoomStats `json:"stats"`
}

type RoomStatusResponse struct {
	Response
	Room Room `json:"room"`
}

type RoomTerrainResponse struct {
	Response
	Terrain []RoomTerrain `json:"terrain"`
}

type SignInResponse struct {
	Response
	Token string `json:"token"`
}

type UpsertResponse struct {
	Response
	Result     Document      `json:"result"`
	Connection Connection    `json:"connection"`
	Message    UpsertMessage `json:"message"`
}

type UserFindResponse struct {
	Response
	User User `json:"user"`
}

type UserOverviewResponse struct {
	Response
	Shards map[string]Shard   `json:"shards"`
	Totals StatsWithIntervals `json:"totals"`
}

type UserRoomsResponse struct {
	Response
	Shards map[string][]string `json:"shards"`
}

type UserStatsResponse struct {
	Response
	Stats StatsWithIntervals `json:"stats"`
}

type VersionResponse struct {
	Response
	Package    int        `json:"package"`
	Protocol   int        `json:"protocol"`
	ServerData ServerData `json:"serverData"`
}

type WorldStartRoomResponse struct {
	Response
	Rooms []string `json:"room"`
}

type WorldStatusResponse struct {
	Response
	Status WorldStatus `json:"status"`
}

type XsollaUserResponse struct {
	Response
	Active int `json:"active"`
}
