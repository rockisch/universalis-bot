package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"universalis-bot/api"
	"universalis-bot/bot"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func runBot() {
	ds, err := bot.GetSession(Token)
	if err != nil {
		panic(err)
	}
	err = ds.Open()
	if err != nil {
		panic(err)
	}

	log.Println("universalis-bot started")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	ds.Close()
}

func testAPI() {
	ctx := context.Background()
	items := []int{37813, 37812}
	data, err := api.GetMarketCurrent(ctx, items)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("test.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	testAPI()
}
