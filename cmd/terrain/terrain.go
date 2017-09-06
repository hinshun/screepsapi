package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hinshun/screepsapi"
)

type Tile uint8

const (
	TilePlain Tile = 1
	TileSwamp      = 2
	TileExit       = 3
	TileWall       = 4
)

type RGBA struct {
	R uint32
	G uint32
	B uint32
	A uint32
}

var (
	roomNames = []string{
		"E40S40",
		"E40S41",
		"E40S42",
		"E40S43",
		"E40S44",
		"E40S45",
		"E40S46",
		"E40S47",
		"E40S48",
		"E40S49",
		"E41S40",
		"E41S41",
		"E41S42",
		"E41S43",
		"E41S44",
		"E41S45",
		"E41S46",
		"E41S47",
		"E41S48",
		"E41S49",
		"E42S40",
		"E42S41",
		"E42S42",
		"E42S43",
		"E42S44",
		"E42S45",
		"E42S46",
		"E42S47",
		"E42S48",
		"E42S49",
		"E43S40",
		"E43S41",
		"E43S42",
		"E43S43",
		"E43S44",
		"E43S45",
		"E43S46",
		"E43S47",
		"E43S48",
		"E43S49",
		"E44S40",
		"E44S41",
		"E44S42",
		"E44S43",
		"E44S44",
		"E44S45",
		"E44S46",
		"E44S47",
		"E44S48",
		"E44S49",
	}

	terrainToTile = map[rune]Tile{
		'0': TilePlain,
		'1': TileWall,
		'2': TileSwamp,
		'3': TileWall,
	}

	rgbaToTile = map[RGBA]Tile{
		RGBA{0x2b2b, 0x2b2b, 0x2b2b, 0xffff}: TilePlain,
		RGBA{0x2323, 0x2525, 0x1313, 0xffff}: TileSwamp,
		RGBA{0, 0, 0, 0xffff}:                TileWall,
		RGBA{0x3232, 0x3232, 0x3232, 0xffff}: TilePlain, // TileExit
	}

	tileToString = map[Tile]string{
		TilePlain: ".",
		TileSwamp: "_",
		TileExit:  "*",
		TileWall:  "#",
	}
)

type Room struct {
	Terrain [][]Tile
}

func NewRoom() *Room {
	terrain := make([][]Tile, 50)
	for i := range terrain {
		terrain[i] = make([]Tile, 50)
	}
	return &Room{
		Terrain: terrain,
	}
}

func (r *Room) Print() {
	for _, row := range r.Terrain {
		var rowStr []string
		for _, rowCell := range row {
			rowStr = append(rowStr, tileToString[rowCell])
		}
		fmt.Println(strings.Join(rowStr, ""))
	}
}

func run() error {
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

	now := time.Now()
	jsonRoom, err := jsonMethod(client, "E48S41")
	if err != nil {
		return err
	}
	fmt.Printf("RoomTerrain: %s\n", time.Since(now))

	now = time.Now()
	pngRoom, err := pngMethod("E48S41")
	if err != nil {
		return err
	}
	fmt.Printf("PNG: %s\n", time.Since(now))

	// now = time.Now()
	// var jsonWG sync.WaitGroup
	// for _, roomName := range roomNames {
	// 	jsonWG.Add(1)
	// 	go func(roomName string) {
	// 		defer jsonWG.Done()
	// 		_, err = jsonMethod(client, roomName)
	// 		if err != nil {
	// 			fmt.Printf("json err: %s\n", err)
	// 		}
	// 	}(roomName)
	// }
	// jsonWG.Wait()
	// fmt.Printf("Parallelized 50x RoomTerrain: %s\n", time.Since(now))

	// now = time.Now()
	// var pngWG sync.WaitGroup
	// for _, roomName := range roomNames {
	// 	pngWG.Add(1)
	// 	go func(roomName string) {
	// 		defer pngWG.Done()
	// 		_, err = pngMethod(roomName)
	// 		if err != nil {
	// 			fmt.Printf("png err: %s\n", err)
	// 		}
	// 	}(roomName)
	// }
	// pngWG.Wait()
	// fmt.Printf("Parallelized 50x PNG: %s\n", time.Since(now))

	fmt.Println("---JSON---")
	jsonRoom.Print()
	fmt.Println("----------")

	fmt.Println("---PNG----")
	pngRoom.Print()
	fmt.Println("----------")

	return nil
}

func jsonMethod(client screepsapi.APIClient, roomName string) (*Room, error) {
	resp, err := client.RoomTerrain("shard1", roomName, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get room terrain json: %s\n", err)
	}

	room := NewRoom()
	for index, char := range resp.Terrain[0].EncodedTerrain {
		tile := terrainToTile[char]
		i := index / 50
		j := index % 50
		room.Terrain[i][j] = tile
	}

	return room, nil
}

func pngMethod(roomName string) (*Room, error) {
	resp, err := http.Get(fmt.Sprintf("https://d3os7yery2usni.cloudfront.net/map/shard1/%s.png", roomName))
	if err != nil {
		return nil, fmt.Errorf("failed to get room terrain png: %s", err)
	}
	defer resp.Body.Close()

	img, err := png.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	room := NewRoom()
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			r, g, b, a := img.At(i*3, j*3).RGBA()
			rgba := RGBA{r, g, b, a}
			room.Terrain[j][i] = rgbaToTile[rgba]
		}
	}

	return room, nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
