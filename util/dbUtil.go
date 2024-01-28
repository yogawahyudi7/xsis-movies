package util

import (
	"fmt"
	"movies-xsis/config"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func DBConnection(config *config.ServerConfig) *gorm.DB {

	set := config.Database

	dsnString := []string{"host=", set.Host, " user=", set.Username, " password=", set.Password, " dbname=", set.DBName, " port=", set.Port, " sslmode=", set.Encrypt, " TimeZone=", set.TimeZone}
	dsn := strings.Join(dsnString, "")

	fmt.Println("--DNS CONNECTION--")
	fmt.Println(dsn)
	fmt.Println("----")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // open connection

	if err != nil {
		panic(err)
	}
	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(20).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db
}
