package database

import (
	"log"
	"os"

	"github.com/Sansui233/proxypool/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connect() (err error) {
	// localhost url
	dsn := "user=proxypool password=proxypool dbname=proxypool port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	if url := config.Config.DatabaseUrl; url != "" {
		dsn = url
	}
	if url := os.Getenv("DATABASE_URL"); url != "" {
		dsn = url
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err == nil {
		log.Println("Database: successfully connected to: ", DB.Name())
	} else {
		log.Println("\n\t\t[db.go] DB connect failed OR no DB. Only use cache to store proxies: ", err)
	}
	return
}
