package utilities

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, _ := sqlmock.New()
	dbConnection, _ := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn:                 db,
				PreferSimpleProtocol: true,
			},
		),
		&gorm.Config{SkipDefaultTransaction: true},
	)

	return mock, dbConnection.Debug()
}
