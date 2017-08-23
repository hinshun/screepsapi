package screepsapi

const (
	// Endpoints
	apiPath = "api"

	// Xsolla endpoints
	xsollaUserPath = "xsolla/user"

	// Auth endpoints
	authMePath     = "auth/me"
	authSigninPath = "auth/signin"

	// Leaderboard endpoints
	leaderboardFindPath    = "leaderboard/find"
	leaderboardListPath    = "leaderboard/list"
	leaderboardSeasonsPath = "leaderboard/seasons"

	// Game endpoints
	addObjectIntentPath       = "game/add-object-intent"
	changeFlagColorPath       = "game/change-flag-color"
	changeFlagPath            = "game/change-flag"
	createConstructionPath    = "game/create-construction"
	createFlagPath            = "game/create-flag"
	genUniqueObjectNamePath   = "game/gen-unique-object-name"
	mapStatsPath              = "game/map-stats"
	marketMyOrdersPath        = "game/market/my-orders"
	marketOrdersIndexPath     = "game/market/orders-index"
	marketOrdersPath          = "game/market/orders"
	roomOverviewPath          = "game/room-overview"
	roomStatusPath            = "game/room-status"
	roomTerrainPath           = "game/room-terrain"
	setNotifyWhenAttackedPath = "game/set-notify-when-attacked"
	timePath                  = "game/time"

	// User endpoints
	userMessagesIndexPath          = "user/messages/index"
	userMessagesListPath           = "user/messages/list"
	userMessagesSendPath           = "user/messages/send"
	userMessagesUnreadCountPath    = "user/messages/unread-count"
	userActivatePTRPath            = "user/activate-ptr"
	userCodePath                   = "user/code"
	userConsolePath                = "user/console"
	userFindPath                   = "user/find"
	userMemoryPath                 = "user/memory"
	userMemorySegmentPath          = "user/memory-segment"
	userMoneyHistoryPath           = "user/money-history"
	userOverviewPath               = "user/overview"
	userRespawnProhibitedRoomsPath = "user/respawn-prohibited-rooms"
	userRoomsPath                  = "user/rooms"
	userStatsPath                  = "user/stats"
	userWorldStartRoomPath         = "user/world-start-room"
	userWorldStatusPath            = "user/world-status"

	// Query params
	encodedKey    = "encoded"
	intervalKey   = "interval"
	limitKey      = "limit"
	modeKey       = "mode"
	offsetKey     = "offset"
	roomKey       = "room"
	seasonKey     = "season"
	shardKey      = "shard"
	usernameKey   = "username"
	respondentKey = "respondent"

	// Query values
	queryTrue = "true"
)
