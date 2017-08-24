package screepsapi

const (
	// Endpoints
	apiPath     = "api"
	versionPath = "version"

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
	// addObjectIntentPath       = "game/add-object-intent"
	// changeFlagColorPath       = "game/change-flag-color"
	// changeFlagPath            = "game/change-flag"
	// createConstructionPath    = "game/create-construction"
	// createFlagPath            = "game/create-flag"
	// genUniqueObjectNamePath   = "game/gen-unique-object-name"
	// mapStatsPath              = "game/map-stats"
	// setNotifyWhenAttackedPath = "game/set-notify-when-attacked"
	timePath = "game/time"

	// Market endpoints
	marketMyOrdersPath    = "game/market/my-orders"
	marketOrdersIndexPath = "game/market/orders-index"
	marketOrdersPath      = "game/market/orders"

	// Messages endpoints
	messagesIndexPath       = "user/messages/index"
	messagesListPath        = "user/messages/list"
	messagesSendPath        = "user/messages/send"
	messagesUnreadCountPath = "user/messages/unread-count"

	// Room endpoints
	roomOverviewPath = "game/room-overview"
	roomStatusPath   = "game/room-status"
	roomTerrainPath  = "game/room-terrain"

	// User endpoints
	// activatePTRPath            = "user/activate-ptr"
	codePath            = "user/code"
	branchesPath        = "user/branches"
	setActiveBranchPath = "user/set-active-branch"
	consolePath         = "user/console"
	memoryPath          = "user/memory"
	// memorySegmentPath          = "user/memory-segment"
	moneyHistoryPath           = "user/money-history"
	respawnProhibitedRoomsPath = "user/respawn-prohibited-rooms"
	userFindPath               = "user/find"
	userOverviewPath           = "user/overview"
	userRoomsPath              = "user/rooms"
	userStatsPath              = "user/stats"

	// World endpoints
	worldStartRoomPath = "user/world-start-room"
	worldStatusPath    = "user/world-status"

	// Socket endpoints
	// TODO(hinshun): no auth required
	socketInfo = "socket/info"

	// Query params
	encodedKey      = "encoded"
	idKey           = "id"
	intervalKey     = "interval"
	limitKey        = "limit"
	modeKey         = "mode"
	offsetKey       = "offset"
	pathKey         = "path"
	resourceTypeKey = "resourceType"
	respondentKey   = "respondent"
	roomKey         = "room"
	seasonKey       = "season"
	shardKey        = "shard"
	statNameKey     = "statName"
	usernameKey     = "username"

	// Query values
	queryTrue = "true"
)
