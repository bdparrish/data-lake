package ingest

import (
	"fmt"
	golog "log"

	"github.com/bufbuild/protovalidate-go"
	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/converter"
	"github.com/codingexplorations/data-lake/common/pkg/db"
	dbModels "github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/proto"
)

type IngestProcessor interface {
	ProcessFolder(folder string) ([]*dbModels.Object, error)
	ProcessFile(fileName string) (*dbModels.Object, error)
}

func GetIngestProcessor(conf *config.Config) IngestProcessor {
	switch conf.IngestProcessorType {
	case "local":
		golog.Println("Using local ingest processor")
		return NewLocalIngestProcessor()
	case "localstack":
		gormDb := db.NewPostgresDb(conf)
		return NewS3IngestProcessorImpl(conf, gormDb)
	default:
		golog.Println("Using default ingest processor")
		return NewLocalIngestProcessor()
	}
}

func validate(object *dbModels.Object) (bool, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return false, fmt.Errorf("failed to initialize proto validator: %v", err)
	}

	protoObject, err := converter.JsonMarshallingConverter[proto.Object, dbModels.Object]{}.DbToProto(object)
	if err != nil {
		return false, fmt.Errorf("failed to convert object: %v", err)
	}

	if err := validator.Validate(protoObject); err != nil {
		return false, fmt.Errorf("failed to validate object: %v", err)
	}

	return true, nil
}
