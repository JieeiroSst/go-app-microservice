package postgres

import (
	"fmt"
	"github.com/JIeeiroSst/go-app/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewPsqlDB(c *config.Config) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPassword,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Connect")
	}
	err = db.AutoMigrate()
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Connect")
	}

	return db, nil
}
