package main

import (
	"audio-storage/config"
	"audio-storage/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}

	app.Run(cfg)

}
