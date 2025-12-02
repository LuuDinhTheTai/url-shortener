package main

import (
	"log"
	"url-shortener/config"
	server2 "url-shortener/internal/server"
)

func main() {
	cfg := config.LoadEnv()
	log.Println(cfg)

	server := server2.NewServer(cfg)

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
