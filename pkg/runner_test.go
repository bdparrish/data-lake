package pkg

import (
	"testing"

	"github.com/codingexplorations/data-lake/models/v1/db"
	"github.com/codingexplorations/data-lake/pkg/config"
	mocks "github.com/codingexplorations/data-lake/test/mocks/pkg/ingest"
)

func TestRunner(t *testing.T) {
	conf := config.GetConfig()
	processor := mocks.NewIngestProcessor(t)

	processor.On("ProcessFolder", "/tmp/data-lake").Return([]*db.Object{}, nil)

	NewRunner(conf, processor).Run()
}
