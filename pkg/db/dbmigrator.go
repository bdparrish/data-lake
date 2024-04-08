package db

import (
	"errors"
	"fmt"

	"github.com/codingexplorations/data-lake/models/v1/db"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrInitialization       = errors.New("error in initializing DB DSN")
	ErrRawConnect           = errors.New("error in connecting to DB (raw)")
	ErrRawMigration         = errors.New("error in running raw migration")
	ErrConstraintsConnect   = errors.New("error in connecting to DB (constraints)")
	ErrConstraintsMigration = errors.New("error in running constraints migration")
)

type Migrator struct {
	dsn    string
	models []interface{}
	logger log.Logger
}

func NewMigrator(logger log.Logger, config *config.Config) (Migrator, error) {
	if config.PostgresHost == "" {
		return Migrator{}, errors.Join(ErrInitialization, fmt.Errorf("postgres host not set"))
	}

	dsn := fmt.Sprintf("host=%s", config.PostgresHost)

	if config.PostgresPort != "" {
		dsn = fmt.Sprintf("%s port=%s", dsn, config.PostgresPort)
	}

	if config.PostgresUser != "" {
		dsn = fmt.Sprintf("%s user=%s", dsn, config.PostgresUser)
	}

	if config.PostgresPassword != "" {
		dsn = fmt.Sprintf("%s password=%s", dsn, config.PostgresPassword)
	}

	if config.PostgresDb != "" {
		dsn = fmt.Sprintf("%s dbname=%s", dsn, config.PostgresDb)
	}

	if config.PostgresSslMode != "" {
		dsn = fmt.Sprintf("%s sslmode=%s", dsn, config.PostgresSslMode)
	}

	return Migrator{
			dsn: dsn,
			models: []interface{}{
				&db.Object{},
			},
			logger: logger,
		},
		nil
}

func (migrator *Migrator) Migrate() error {
	// first run without constraint creation
	gormDB, err := migrator.connectRaw()
	if err != nil {
		return errors.Join(ErrRawConnect, err)
	}

	migrator.logger.Info("Running Raw Auto Migration")
	if err := gormDB.Debug().AutoMigrate(migrator.models...); err != nil {
		return errors.Join(ErrRawMigration, err)
	}

	// re-run auto-migration with constraints enabled to create explicit data constraints as described in the models
	gormDb, err := migrator.connect()
	if err != nil {
		return errors.Join(ErrConstraintsConnect, err)
	}
	migrator.logger.Info("Running Auto Migration")
	if err := gormDb.Debug().AutoMigrate(migrator.models...); err != nil {
		return errors.Join(ErrConstraintsMigration, err)
	}

	migrator.logger.Info("Auto Migration completed successfully.")
	return nil
}

func (migrator *Migrator) connectRaw() (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(migrator.dsn),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
}

func (migrator *Migrator) connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(migrator.dsn), &gorm.Config{})
}
