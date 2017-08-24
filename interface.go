package screepsapi

type APIClient interface {
	// Auth
	Me(shard string) (MeResponse, error)
	SignIn(email, password string) error

	// Code
	Branches() (BranchesResponse, error)
	Console(shard, expression string) (InsertResponse, error)
	PullCode(branch string) (PullCodeResponse, error)
	PushCode(branch string, modules map[string]string) (PushCodeResponse, error)
	SetActiveBranch(branch string, activeName ActiveName) (Response, error)

	// Game
	AddObjectIntent(addObjectIntentReq AddObjectIntentRequest) (UpsertResponse, error)
	CheckUniqueObjectName(shard, name string, uniqueObjectNameType UniqueObjectName) (Response, error)
	CreateConstruction(createConstructionReq CreateConstructionRequest) (UpsertResponse, error)
	CreateFlag(createFlagReq CreateFlagRequest) (UpsertResponse, error)
	DestroyStructure(shard, room, user string, ids []string) (UpsertResponse, error)
	GenUniqueObjectName(shard string, uniqueObjectNameType UniqueObjectName) (GenUniqueObjectNameResponse, error)
	PlaceSpawn(placeSpawnReq PlaceSpawnRequest) (PlaceSpawnResponse, error)
	RemoveFlag(shard, room, name string) (UpsertResponse, error)
	RemoveObject(shard, room, id string) (UpsertResponse, error)
	Respawn() (RespawnResponse, error)
	RespawnProhibitedRooms() (RespawnProhibitedRoomsResponse, error)
	SuicideCreep(shard, room, id string) (UpsertResponse, error)
	Time(shard string) (Response, error)
	UnclaimController(shard, room, id string) (UpsertResponse, error)

	// Leaderboard
	LeaderboardFind(mode, username, season string) (LeaderboardFindResponse, error)
	LeaderboardList(mode, season string, limit, offset int) (LeaderboardListResponse, error)
	LeaderboardSeasons() (LeaderboardSeasonsResponse, error)

	// Map
	MapStats(shard string, rooms []string, statName StatName, statsPeriod StatsPeriod) (MapStatsResponse, error)

	// Market
	MoneyHistory() (MoneyHistoryResponse, error)
	Orders(shard, resourceType string) (OrdersResponse, error)
	OrdersIndex(shard string) (OrdersIndexResponse, error)
	MyOrders() (OrdersResponse, error)

	// Memory
	Memory(shard, path string) (MemoryResponse, error)
	UpdateMemory(shard, path, value string) (InsertResponse, error)
	DeleteMemory(shard, path string) (InsertResponse, error)

	// Messages
	MessagesIndex() (MessagesIndexResponse, error)
	MessagesList(respondent string) (MessagesListResponse, error)
	MessagesSend(respondent, text string) (Response, error)
	MessagesUnreadCount() (MessagesUnreadCountResponse, error)

	// Room
	RoomOverview(shard, room string, statsPeriod StatsPeriod) (RoomOverviewResponse, error)
	RoomStatus(shard, room string) (RoomStatusResponse, error)
	RoomTerrain(shard, room string, encoded bool) (RoomTerrainResponse, error)

	// User
	UserFind(id, username string) (UserFindResponse, error)
	UserOverview(statName StatName, statsPeriod StatsPeriod) (UserOverviewResponse, error)
	UserRooms() (UserRoomsResponse, error)
	UserStats(id string, statsPeriod StatsPeriod) (UserStatsResponse, error)

	// Version
	Version() (VersionResponse, error)

	// World
	WorldStartRoom() (WorldStartRoomResponse, error)
	WorldStatus() (WorldStatusResponse, error)

	// Xsolla
	XsollaUser(shard string) (XsollaUserResponse, error)
}
