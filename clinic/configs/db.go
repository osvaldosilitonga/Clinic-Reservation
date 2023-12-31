package configs

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DB_PG_STRING")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError:       true,
		FullSaveAssociations: true,
	})
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	// DB connection pool
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(90)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	log.Println("Database connected")

	return db
}
