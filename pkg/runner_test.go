package pkg

import (
	"testing"

	"github.com/codeexplorations/data-lake/config"
	models_v1 "github.com/codeexplorations/data-lake/models/v1"
	mocks "github.com/codeexplorations/data-lake/test/mocks/pkg/ingest"
)

func TestRunner(t *testing.T) {
	conf := config.GetConfig()
	processor := mocks.NewIngestProcessor(t)

	processor.On("ProcessFolder", "/tmp/data-lake").Return([]*models_v1.Object{}, nil)

	NewRunner(conf, processor).Run()
}
