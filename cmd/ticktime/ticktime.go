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
	"github.com/hinshun/screepsapi/screepsws"
)

func measureTickTime() error {
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

	ws, err := screepsws.NewWebSocket(credentials.ServerURL, client.Token)
	if err != nil {
		return err
	}

	moneyChan, err := ws.SubscribeMoney("599bc57078ca755b8407aa4f")
	if err != nil {
		return err
	}

	go func() {
		var sum int64
		var tick int64
		skip := 10
		for {
			now := time.Now()
			<-moneyChan
			diff := time.Since(now)
			if skip > 0 {
				skip--
			} else {
				tick++
				sum += int64(diff)
				fmt.Printf("avg %s\n", time.Duration(sum/tick))
			}
		}
	}()

	time.Sleep(1 * time.Minute)

	err = ws.UnsubscribeMoney("599bc57078ca755b8407aa4f")
	if err != nil {
		return err
	}

	err = ws.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := measureTickTime()
	if err != nil {
		log.Fatal(err)
	}
}
