package ingest

import (
	"testing"

	dbModels "github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	"github.com/stretchr/testify/assert"
)

func TestFolderIngest_ProcessFile_validate(t *testing.T) {
	object := &dbModels.Object{
		FileName:     "test.txt",
		FileLocation: "/tmp/test/test.txt",
		ContentType:  "text/plain",
		ContentSize:  15,
	}

	valid, err := validate(object)

	assert.Nil(t, err)
	assert.True(t, valid)
}

func TestFolderIngest_ProcessFile_validateTable(t *testing.T) {
	tests := []struct {
		name          string
		object        *dbModels.Object
		expectedError string
	}{
		{
			name: "invalid - FileName empty",
			object: &dbModels.Object{
				FileName:     "",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "text/plain",
				ContentSize:  15,
			},
			expectedError: "file_name: value is required [required]",
		},
		{
			name: "invalid - FileLocation empty",
			object: &dbModels.Object{
				FileName:     "test.txt",
				FileLocation: "",
				ContentType:  "text/plain",
				ContentSize:  15,
			},
			expectedError: "file_location: value is required [required]",
		},
		{
			name: "invalid - ContentType empty",
			object: &dbModels.Object{
				FileName:     "test.txt",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "",
				ContentSize:  15,
			},
			expectedError: "content_type: value is required [required]",
		},
		{
			name: "invalid - ContentSize less than 0",
			object: &dbModels.Object{
				FileName:     "test.txt",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "text/plain",
				ContentSize:  0,
			},
			expectedError: "content_size: value must be greater than 0 and less than or equal to 1048576 [int32.gt_lte]",
		},
		{
			name: "invalid - ContentSize greater than 1GB",
			object: &dbModels.Object{
				FileName:     "test.txt",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "text/plain",
				ContentSize:  2097152,
			},
			expectedError: "content_size: value must be greater than 0 and less than or equal to 1048576 [int32.gt_lte]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := validate(tc.object)

			assert.Error(t, err)
			assert.False(t, valid)
			assert.Contains(t, err.Error(), tc.expectedError)
		})
	}
}
