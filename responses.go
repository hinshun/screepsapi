package screepsapi

import (
	"github.com/hinshun/screepsapi/screepstype"
)

type APIResponse interface {
	IsOk() bool
}

type Response struct {
	Ok    int               `json:"ok"`
	Error screepstype.Error `json:"error"`
}

func (r *Response) IsOk() bool {
	return r.Ok == 1
}

type BranchesResponse struct {
	Response
	Branches []screepstype.Branch `json:"list"`
}

type GenUniqueObjectNameResponse struct {
	Response
	Name string `json:"name"`
}

type InsertResponse struct {
	Response
	Result        screepstype.InsertResult `json:"result"`
	Ops           []screepstype.InsertOps  `json:"ops"`
	InsertedCount int                      `json:"insertedCount"`
	InsertedIDs   []string                 `json:"insertedIDs"`
}

type LeaderboardFindResponse struct {
	Response
	screepstype.LeaderboardPosition
	Positions []screepstype.LeaderboardPosition `json:"list"`
}

type LeaderboardListResponse struct {
	Response
	Positions []screepstype.LeaderboardPosition `json:"list"`
	Count     int                               `json:"count"`
	Users     map[string]screepstype.User       `json:"users"`
}

type LeaderboardSeasonsResponse struct {
	Response
	Seasons []screepstype.Season `json:"seasons"`
}

type MapStatsResponse struct {
	Response
	GameTime int                                 `json:"gameTime"`
	Stats    map[string]screepstype.MapRoomStats `json:"stats"`
	Users    map[string]screepstype.User         `json:"users"`
}

type MeResponse struct {
	Response
	ID                 string             `json:"_id"`
	Email              string             `json:"email"`
	Username           string             `json:"username"`
	CPU                int                `json:"cpu"`
	Badge              screepstype.Badge  `json:"badge"`
	Password           bool               `json:"password"`
	Credits            int                `json:"credits"`
	LastRespawnDate    int                `json:"lastRespawnDate"`
	PromoPeriodUntil   int                `json:"promoPeriodUntil"`
	Money              int                `json:"money"`
	SubscriptionTokens int                `json:"subscriptionTokens"`
	GitHub             screepstype.GitHub `json:"github"`
	Steam              screepstype.Steam  `json:"steam"`
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
	Messages []screepstype.MessageWithRecipients `json:"messages"`
	Users    map[string]screepstype.User         `json:"users"`
}

type MessagesListResponse struct {
	Response
	Messages []screepstype.Message `json:"messages"`
}

type MessagesUnreadCountResponse struct {
	Response
	Count int `json:"count"`
}

type MoneyHistoryResponse struct {
	Response
	Page     int                        `json:"page"`
	HasMore  bool                       `json:"hasMore"`
	Messages []screepstype.MoneyMessage `json:"messages"`
}

type MyOrdersResponse struct {
	Response
	Orders []screepstype.MarketOrder `json:"list"`
}

type OrdersIndexResponse struct {
	Response
	ResourceCounts []screepstype.ResourceCount `json:"list"`
}

type OrdersResponse struct {
	Response
	Orders []screepstype.MarketOrder `json:"list"`
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
	First  screepstype.Document `json:"0"`
	Second screepstype.Document `json:"1"`
}

type RespawnProhibitedRoomsResponse struct {
	Response
	Rooms []string `json:"rooms"`
}

type RoomOverviewResponse struct {
	Response
	Owner screepstype.RoomOwner `json:"owner"`
	Stats screepstype.RoomStats `json:"stats"`
}

type RoomStatusResponse struct {
	Response
	Room screepstype.Room `json:"room"`
}

type RoomTerrainResponse struct {
	Response
	Terrain []screepstype.RoomTerrain `json:"terrain"`
}

type SignInResponse struct {
	Response
	Token string `json:"token"`
}

type TimeResponse struct {
	Response
	Time int `json:"time"`
}

type UpsertResponse struct {
	Response
	Result     screepstype.Document      `json:"result"`
	Connection screepstype.Connection    `json:"connection"`
	Message    screepstype.UpsertMessage `json:"message"`
}

type UserFindResponse struct {
	Response
	User screepstype.User `json:"user"`
}

type UserOverviewResponse struct {
	Response
	Shards map[string]screepstype.Shard   `json:"shards"`
	Totals screepstype.StatsWithIntervals `json:"totals"`
}

type UserRoomsResponse struct {
	Response
	Shards map[string][]string `json:"shards"`
}

type UserStatsResponse struct {
	Response
	Stats screepstype.StatsWithIntervals `json:"stats"`
}

type VersionResponse struct {
	Response
	Package    int                    `json:"package"`
	Protocol   int                    `json:"protocol"`
	ServerData screepstype.ServerData `json:"serverData"`
}

type WorldStartRoomResponse struct {
	Response
	Rooms []string `json:"room"`
}

type WorldStatusResponse struct {
	Response
	Status screepstype.WorldStatus `json:"status"`
}

type XsollaUserResponse struct {
	Response
	Active int `json:"active"`
}
