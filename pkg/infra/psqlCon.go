package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConn() (res *gorm.DB) {
	// you can replace the below string by parsing env variables
	dsn := "host=localhost user=postgres password=postgres dbname=service_user port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	return db
}
