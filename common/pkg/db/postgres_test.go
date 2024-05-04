package db

import (
	"testing"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewPostgresDb(t *testing.T) {
	conf := config.GetConfig()

	db := NewPostgresDb(conf)

	assert.NotNil(t, db)
}
