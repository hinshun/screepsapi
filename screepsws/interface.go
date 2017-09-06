package screepsws

type WebSocket interface {
	Close() error

	// Server
	SubscribeServerMessage() (<-chan ServerMessageResponse, error)
	UnsubscribeServerMessage() error

	// Room
	SubscribeRoom(shard, room string) (<-chan RoomResponse, error)
	UnsubscribeRoom(shard, room string) error
	SubscribeRoomMap(shard, room string) (<-chan RoomMapResponse, error)
	UnsubscribeRoomMap(shard, room string) error

	// User
	SubscribeCode(userID string) (<-chan CodeResponse, error)
	UnsubscribeCode(userID string) error
	SubscribeConsole(userID string) (<-chan ConsoleResponse, error)
	UnsubscribeConsole(userID string) error
	SubscribeCPU(userID string) (<-chan CPUResponse, error)
	UnsubscribeCPU(userID string) error
	SubscribeMemory(userID, path string) (<-chan MemoryResponse, error)
	UnsubscribeMemory(userID, path string) error
	SubscribeMessage(userID, respondentID string) (<-chan MessageResponse, error)
	UnsubscribeMessage(userID, respondentID string) error
	SubscribeMoney(userID string) (<-chan int, error)
	UnsubscribeMoney(userID string) error
	SubscribeNewMessage(userID string) (<-chan NewMessageResponse, error)
	UnsubscribeNewMessage(userID string) error
	SubscribeSetActiveBranch(userID string) (<-chan SetActiveBranchResponse, error)
	UnsubscribeSetActiveBranch(userID string) error
}
