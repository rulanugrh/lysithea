package config

import (
	"fmt"
	"log"

	"github.com/rulanugrh/lysithea/internal/entity/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	conf := GetConfig()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Jakarta",
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, web.InternalServerError("Cant connect to postgresql")
	}

	log.Print("Success connect to database")

	return db, nil
}
