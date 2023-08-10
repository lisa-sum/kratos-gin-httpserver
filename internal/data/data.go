package data

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tiktok/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPostgres, NewUserRepo)

// Database .
type Database struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Database, logger log.Logger, db *gorm.DB) (*Database, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Database{db: db}, cleanup, nil
}

func NewPostgres(c *conf.Database) *gorm.DB {
	host := c.Postgres.Host
	user := c.Postgres.User
	password := c.Postgres.Password
	dbname := c.Postgres.Dbname
	port := c.Postgres.Port
	sslMode := c.Postgres.SslMode
	timeZone := c.Postgres.TimeZone

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslMode, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
