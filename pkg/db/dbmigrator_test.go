package db

import (
	"reflect"
	"testing"

	"github.com/codingexplorations/data-lake/models/v1/db"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestSettingAllConfig(t *testing.T) {
	logger, _ := log.GetLogger()

	migrator, err := NewMigrator(logger, config.GetConfig())
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "host=postgres port=5432 user=postgres password=testpassword dbname=data_lake sslmode=disable", migrator.dsn)

	expected := []interface{}{
		&db.Object{},
	}
	assert.True(t, reflect.DeepEqual(expected, migrator.models))
}
