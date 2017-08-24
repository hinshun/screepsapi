package screepsapi

const (
	// Endpoints
	apiPath     = "api"
	versionPath = "version"

	// Xsolla endpoints
	xsollaUserPath = "xsolla/user"

	// Auth endpoints
	mePath     = "auth/me"
	signInPath = "auth/signin"

	// Leaderboard endpoints
	leaderboardFindPath    = "leaderboard/find"
	leaderboardListPath    = "leaderboard/list"
	leaderboardSeasonsPath = "leaderboard/seasons"

	// Game endpoints
	addObjectIntentPath       = "game/add-object-intent"
	checkUniqueObjectNamePath = "game/check-unique-object-name"
	createConstructionPath    = "game/create-construction"
	createFlagPath            = "game/create-flag"
	genUniqueObjectNamePath   = "game/gen-unique-object-name"
	mapStatsPath              = "game/map-stats"
	placeSpawnPath            = "game/place-spawn"
	removeFlagPath            = "game/remove-flag"
	setNotifyWhenAttackedPath = "game/set-notify-when-attacked"
	timePath                  = "game/time"

	// Market endpoints
	myOrdersPath    = "game/market/my-orders"
	ordersIndexPath = "game/market/orders-index"
	ordersPath      = "game/market/orders"

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
	branchesPath = "user/branches"
	codePath     = "user/code"
	consolePath  = "user/console"
	memoryPath   = "user/memory"
	// memorySegmentPath          = "user/memory-segment"
	// TODO(hinshun): unconfirmed struct types
	moneyHistoryPath           = "user/money-history"
	respawnPath                = "user/respawn"
	respawnProhibitedRoomsPath = "user/respawn-prohibited-rooms"
	setActiveBranchPath        = "user/set-active-branch"
	userFindPath               = "user/find"
	userOverviewPath           = "user/overview"
	// TODO(hinshun): unconfirmed struct types and query params
	userRoomsPath = "user/rooms"
	userStatsPath = "user/stats"

	// World endpoints
	worldStartRoomPath = "user/world-start-room"
	worldStatusPath    = "user/world-status"

	// Socket endpoints
	// TODO(hinshun): no auth required
	socketInfo = "socket/info"

	// Query params
	branchKey       = "branch"
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
