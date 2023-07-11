package main

import (
	"encoding/json"
	"fmt"
	"httpclient/coincap/coin-cap"
	"log"
	"time"
)

func main() {

	client, err := coin_cap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	r, err := client.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	humanReadableResp, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(humanReadableResp))
}
