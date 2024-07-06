// Package postgres returns gorm db
package postgres

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB returns gorm database
func NewDB(dsn string) *gorm.DB {
	var errs error
	for i := 0; i < 5; i++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			errs = err
			time.Sleep(3 * time.Second)
			continue
		}
		log.Println("Database Connected")
		return db
	}
	panic(errs)
}
