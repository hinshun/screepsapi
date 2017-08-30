package screepsws

const (
	// Endpoints
	socketPath = "socket/websocket"

	// Messages
	authFormat        = "auth %s"
	gzipFormat        = "gzip %s"
	subscribeFormat   = "subscribe %s"
	unsubscribeFormat = "unsubscribe %s"

	// Global subscriptions
	// TODO(hinshun): unimplemented
	serverMessageFormat = "server-message"

	// User subscriptions
	codeFormat    = "user:%s/code"
	consoleFormat = "user:%s/console"
	cpuFormat     = "user:%s/cpu"
	// TODO(hinshun): unimplemented
	memoryFormat = "user:%s/memory/%s"
	// TODO(hinshun): unimplemented
	messageFormat = "user:%s/message:%s"
	moneyFormat   = "user:%s/money"
	// TODO(hinshun): unimplemented
	newMessageFormat      = "user:%s/newMessage"
	setActiveBranchFormat = "user:%s/set-active-branch"

	// Room subscriptions
	roomFormat    = "room:%s/%s"
	roomMapFormat = "roomMap2:%s/%s"
)
