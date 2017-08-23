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

	// marketOrdersIndex, err := client.MarketOrdersIndex("shard1")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("market/orders-index: %#v\n", marketOrdersIndex)

	// userMessagesIndex, err := client.UserMessagesIndex()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/index: %#v\n", userMessagesIndex)

	// userMessagesList, err := client.UserMessagesList("59947fba7c2c8c7cccccfc41")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/list: %#v\n", userMessagesList)

	// userMessagesSend, err := client.UserMessagesSend("59947fba7c2c8c7cccccfc41", "nvm")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/send: %#v\n", userMessagesSend)

	// userMessagesUnreadCount, err := client.UserMessagesUnreadCount()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/unread-count: %#v\n", userMessagesUnreadCount)

	// userConsole, err := client.UserConsole("shard1", "Game")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/send: %#v\n", userConsole)

	// userFind, err := client.UserFind("", "hinshun")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/find: %#v\n", userFind)

	// userMemory, err := client.UserMemory("shard1", "flags")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", userMemory)

	// userMemoryConsole, err := client.UpdateUserMemory("shard1", "flags", "test")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", userMemoryConsole)

	// userMemoryConsole, err := client.DeleteUserMemory("shard1", "flags")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", userMemoryConsole)

	// TODO(hinshun): unconfirmed struct types
	// userMoneyHistory, err := client.UserMoneyHistory()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/money-history: %#v\n", userMoneyHistory)

	// userOverview, err := client.UserOverview("", screepsapi.StatsPeriodNone)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/overview: %#v\n", userOverview)

	return nil
}

func main() {
	err := test()
	if err != nil {
		log.Fatal(err)
	}
}
