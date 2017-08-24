package main

import (
	"fmt"
	"log"

	"github.com/hinshun/screepsapi"
)

func test() error {
	client, err := screepsapi.NewClient(screepsapi.Credentials{
		Email:     "edgarhinshunlee@gmail.com",
		Password:  "X9*pBPYS3WgywbR12EgqcjKgoC2#RJZh",
		ServerURL: "https://screeps.com",
	})
	if err != nil {
		return fmt.Errorf("failed to create screepsapi client: %s", err)
	}

	// xsollaUser, err := client.XsollaUser("shard1")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("xsolla/user: %#v\n", xsollaUser)

	// authMe, err := client.AuthMe("shard1")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("auth/me: %#v\n", authMe)

	// leaderboardFind, err := client.LeaderboardFind("power", "Dissi", "")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("leaderboard/find: %#v\n", leaderboardFind)

	// leaderboardList, err := client.LeaderboardList("power", "2017-08", 10, 0)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("leaderboard/list: %#v\n", leaderboardList)

	// leaderboardSeasons, err := client.LeaderboardSeasons()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("leaderboard/seasons: %#v\n", leaderboardSeasons)

	// roomOverview, err := client.RoomOverview("shard1", "E23S26", screepsapi.StatsPeriod1Hour)
	// if err != nil {
	//	return err
	// }
	// fmt.Printf("game/room-overview: %#v\n", roomOverview)

	// roomTerrain, err := client.RoomTerrain("shard1", "E23S26", true)
	// if err != nil {
	//	return err
	// }
	// fmt.Printf("game/room-terrain: %#v\n", roomTerrain)

	// roomStatus, err := client.RoomStatus("shard1", "E47S16")
	// if err != nil {
	//	return err
	// }
	// fmt.Printf("game/room-status: %#v\n", roomStatus)

	// marketMyOrders, err := client.MarketMyOrders()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("market/my-orders: %#v\n", marketMyOrders)

	// marketOrdersIndex, err := client.MarketOrdersIndex("shard1")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("market/orders-index: %#v\n", marketOrdersIndex)

	marketOrders, err := client.MarketOrders("shard1", "Z")
	if err != nil {
		return err
	}
	fmt.Printf("market/orders: %#v\n", marketOrders)

	// time, err := client.Time("shard1")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/time: %#v\n", time)

	// messagesIndex, err := client.MessagesIndex()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/index: %#v\n", messagesIndex)

	// messagesList, err := client.MessagesList("59947fba7c2c8c7cccccfc41")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/list: %#v\n", messagesList)

	// messagesSend, err := client.MessagesSend("59947fba7c2c8c7cccccfc41", "nvm")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/send: %#v\n", messagesSend)

	// messagesUnreadCount, err := client.MessagesUnreadCount()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/unread-count: %#v\n", messagesUnreadCount)

	// console, err := client.Console("shard1", "Game")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/send: %#v\n", console)

	// userFind, err := client.UserFind("", "hinshun")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/find: %#v\n", userFind)

	// memory, err := client.Memory("shard1", "flags")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", memory)

	// memoryConsole, err := client.UpdateMemory("shard1", "flags", "test")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", memoryConsole)

	// memoryConsole, err := client.DeleteMemory("shard1", "flags")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", memoryConsole)

	// TODO(hinshun): unconfirmed struct types
	// moneyHistory, err := client.MoneyHistory()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/money-history: %#v\n", moneyHistory)

	// userOverview, err := client.UserOverview("", screepsapi.StatsPeriodNone)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/overview: %#v\n", userOverview)

	// respawnProhibitedRooms, err := client.RespawnProhibitedRooms()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/respawn-prohibited-rooms: %#v\n", respawnProhibitedRooms)

	// TODO(hinshun): unconfirmed struct types and query params
	// userRooms, err := client.UserRooms()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/rooms: %#v\n", userRooms)

	// userStats, err := client.UserStats("561e4d4645f3f7244a7622e8", screepsapi.StatsPeriod7Days)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/stats: %#v\n", userStats)

	// worldStartRoom, err := client.WorldStartRoom()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/world-start-room: %#v\n", worldStartRoom)

	// worldStatus, err := client.WorldStatus()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/world-status: %#v\n", worldStatus)

	return nil
}

func main() {
	err := test()
	if err != nil {
		log.Fatal(err)
	}
}
