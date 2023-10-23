package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConn() (res *gorm.DB) {
	dsn := "host=localhost user=postgres password=postgres dbname=service_user port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	return db
}
