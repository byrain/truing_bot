package main

import (
	"fmt"

	turing "github.com/byrain/turing_bot/turing"
)

func main() {
	message := turing.NewTuringMessage("北京天气")
	messageResp := turing.GetTuringBotResp(message)
	fmt.Println(messageResp.Result[0].Values["text"])
}
