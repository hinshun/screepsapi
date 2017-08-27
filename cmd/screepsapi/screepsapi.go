package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/hinshun/screepsapi"
)

func test() error {
	credentialsFile, err := os.Open("credentials.json")
	if err != nil {
		return fmt.Errorf("failed to open credentials file: %s", err)
	}

	credentialsBytes, err := ioutil.ReadAll(bufio.NewReader(credentialsFile))
	if err != nil {
		return fmt.Errorf("failed to read credentials file: %s", err)
	}

	credentials := screepsapi.Credentials{}
	err = json.Unmarshal(credentialsBytes, &credentials)
	if err != nil {
		return fmt.Errorf("failed to unmarshal credentials file: %s", err)
	}

	client, err := screepsapi.NewClient(credentials)
	if err != nil {
		return fmt.Errorf("failed to create screepsapi client: %s", err)
	}

	// version, err := client.Version()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("version: %#v\n", version)

	err = client.WebSocket.Connect()
	if err != nil {
		return err
	}

	cpuChan, err := client.WebSocket.SubscribeCPU("599bc57078ca755b8407aa4f")
	if err != nil {
		return err
	}

	go func() {
		for {
			cpuData := <-cpuChan
			fmt.Printf("cpu-data: %#v\n", cpuData)
		}
	}()

	time.Sleep(10 * time.Second)

	err = client.WebSocket.Close()
	if err != nil {
		return err
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

	// marketOrders, err := client.MarketOrders("shard1", "Z")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("market/orders: %#v\n", marketOrders)

	// placeSpawn, err := client.PlaceSpawn(screepsapi.PlaceSpawnRequest{
	// 	Shard:     "shard1",
	// 	Room:      "E39N17",
	// 	X:         27,
	// 	Y:         24,
	// 	Name: "Spawn1",
	// })
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/place-spawn: %#v\n", placeSpawn)

	// removeFlag, err := client.RemoveFlag("shard1", "E7N35", "Flag1")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/remove-flag: %#v\n", removeFlag)

	// removeObject, err := client.RemoveObject("shard1", "E29N39", "599e8f01368c125b1639e407")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/add-object-intent: %#v\n", removeObject)

	// suicideCreep, err := client.SuicideCreep("shard1", "E29N39", "599e92c84e56f12ea0b06109")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/add-object-intent: %#v\n", suicideCreep)

	// destroyStructure, err := client.DestroyStructure("shard1", "E39N17", "599bc57078ca755b8407aa4c", []string{"599eea6cc29e4d5b236775a8"})
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/add-object-intent: %#v\n", destroyStructure)

	// createConstruction, err := client.CreateConstruction(screepsapi.CreateConstructionRequest{
	// 	Shard:     "shard1",
	// 	Room:      "E39N17",
	// 	X:         27,
	// 	Y:         24,
	// 	Structure: screepsapi.StructureSpawn,
	// })
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/create-construction: %#v\n", createConstruction)

	// createFlag, err := client.CreateFlag(screepsapi.CreateFlagRequest{
	// 	Shard:          "shard1",
	// 	Room:           "E7N35",
	// 	X:              19,
	// 	Y:              37,
	// 	Name:           "test",
	// 	Color:          screepsapi.ColorWhite,
	// 	SecondaryColor: screepsapi.ColorRed,
	// })
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/create-flag: %#v\n", createFlag)

	// genUniqueObjectName, err := client.GenUniqueObjectName("shard1", screepsapi.UniqueObjectNameSpawn)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/gen-unique-object-name: %#v\n", genUniqueObjectName)

	// checkUniqueObjectName, err := client.CheckUniqueObjectName("shard1", "Spawn1", screepsapi.UniqueObjectNameSpawn)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/check-unique-object-name: %#v\n", checkUniqueObjectName)

	// mapStats, err := client.MapStats("shard1", []string{"E7N35"}, screepsapi.StatNameEnergyHarvested, screepsapi.StatsPeriod1Hour)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("game/map-stats: %#v\n", mapStats)

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

	// branches, err := client.Branches()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/branches: %#v\n", branches)

	// pushCode, err := client.PushCode("tutorial-1", map[string]string{
	// 	"main": "hello world",
	// })
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/code: %#v\n", pushCode)

	// pullCode, err := client.PullCode("")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/code: %#v\n", pullCode)

	// insert, err := client.Console("shard1", "Game")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/messages/send: %#v\n", insert)

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

	// memoryUpdate, err := client.UpdateMemory("shard1", "flags", "test")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", memoryUpdate)

	// memoryDelete, err := client.DeleteMemory("shard1", "flags")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory: %#v\n", memoryDelete)

	// memorySegment, err := client.MemorySegment("shard1", 0)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory-segment: %#v\n", memorySegment)

	// updateMemorySegment, err := client.UpdateMemorySegment("shard1", "", 0)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/memory-segment: %#v\n", updateMemorySegment)

	// moneyHistory, err := client.MoneyHistory()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/money-history: %#v\n", moneyHistory)

	// userOverview, err := client.UserOverview(screepsapi.StatNameNone, screepsapi.StatsPeriodNone)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/overview: %#v\n", userOverview)

	// respawn, err := client.Respawn()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/respawn: %#v\n", respawn)

	// respawnProhibitedRooms, err := client.RespawnProhibitedRooms()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/respawn-prohibited-rooms: %#v\n", respawnProhibitedRooms)

	// setActiveBranch, err := client.SetActiveBranch("world", screepsapi.ActiveNameWorld)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("user/set-active-branch: %#v\n", setActiveBranch)

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
