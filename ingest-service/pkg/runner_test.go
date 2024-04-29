package pkg

import (
	"testing"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	mocks "github.com/codingexplorations/data-lake/ingest-service/test/mocks/pkg/ingest"
)

func TestRunner(t *testing.T) {
	conf := config.GetConfig()
	processor := mocks.NewIngestProcessor(t)

	processor.On("ProcessFolder", "/tmp/data-lake").Return([]*db.Object{}, nil)

	NewRunner(conf, processor).Run()
}
