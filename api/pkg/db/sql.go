package db

import (
	"fmt"
	"os"

	"github.com/mixedmachine/GoalsBackend/api/pkg/models"
	"gorm.io/driver/postgres"

	"log"

	"gorm.io/gorm"
)

type SqlConnection interface {
	Close()
	DB() *gorm.DB
}

type sqlConn struct {
	client  *gorm.DB
	session gorm.DB
}

func NewSqlConnection() SqlConnection {
	var c sqlConn
	var err error

	db := postgres.Open(getURL())

	c.client, err = gorm.Open(db, &gorm.Config{})
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *sqlConn) Close() {

}

func (c *sqlConn) DB() *gorm.DB {

	c.client = c.client.Session(&gorm.Session{
		// TODO: Add additional session configs
		CreateBatchSize: 1000,
	})
	db, err := c.client.DB()
	if err != nil {
		log.Panicln(err.Error())
	}
	if err := db.Ping(); err != nil {
		return nil
	}
	err = c.client.AutoMigrate(models.Goal{})
	if err != nil {
		return nil
	}
	return c.client
}

func getURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
}
