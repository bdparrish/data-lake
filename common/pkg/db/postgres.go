package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDb(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDb,
		config.PostgresPort,
		config.PostgresSslMode,
	)

	pgDB, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	pgDB.SetConnMaxLifetime(time.Duration(config.PostgresMaxConnTimeMinutes) * time.Minute)
	pgDB.SetMaxIdleConns(config.PostgresMaxIdleConnections)
	pgDB.SetMaxOpenConns(config.PostgresMaxOpenConnections)

	database, err := gorm.Open(
		postgres.New(postgres.Config{Conn: pgDB}),
		&gorm.Config{TranslateError: true},
	)
	if err != nil {
		panic(err)
	}

	return database
}
