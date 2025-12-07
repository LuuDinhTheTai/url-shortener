package main

import (
	"log"
	"url-shortener/config"
	"url-shortener/internal/model"
	server2 "url-shortener/internal/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadEnv()
	log.Println(cfg)

	pgDB, errDB := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.Database.URI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if errDB != nil {
		log.Fatal(errDB)
	}

	errMg := pgDB.AutoMigrate(&model.Url{})
	if errMg != nil {
		log.Fatal(errMg)
	}

	server := server2.NewServer(cfg, pgDB)

	errRun := server.Run()
	if errRun != nil {
		log.Fatal(errRun)
	}
}
