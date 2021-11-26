package main

import (
	"github.com/Liki4/KaiheilaBot/internal/bot"
	"github.com/Liki4/KaiheilaBot/internal/conf"
	"github.com/Liki4/KaiheilaBot/internal/handlers/ncm"
	log "unknwon.dev/clog/v2"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic(err)
	}

	if err = conf.Load(); err != nil {
		log.Fatal("Failed to load config: %v", err)
	}

	ncm.Init()
	bot.Run()
}
