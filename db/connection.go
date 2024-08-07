package db

import (
	"fmt"
	"log"

	"github.com/itzaddddd/ticket-reserve/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConn(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s
	sslmode=disable TimeZone=Asia/Bangkok`, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User,
		cfg.Postgres.Password, cfg.Postgres.DbName)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	postgresDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := postgresDb.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("connect to postgresql success")

	return db

}
