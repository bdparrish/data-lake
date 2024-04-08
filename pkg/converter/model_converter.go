package converter

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/samber/lo"
)

type GenericConverter[proto interface{}, db interface{}] interface {
	DbToProto(*db) (*proto, error)
	ProtoToDb(*proto) (*db, error)
	DbToProtoSlice([]*db) ([]*proto, error)
	ProtoToDbSlice([]*proto) ([]*db, error)
}

type JsonMarshallingConverter[proto interface{}, db interface{}] struct{}

func (converter JsonMarshallingConverter[proto, db]) DbToApi(source *db) (*proto, error) {
	var response proto
	sourceBytes, err := json.Marshal(&source)
	if err != nil {
		return nil, errors.Join(errors.New("error in serializing db model"), err)
	}
	err = json.Unmarshal(sourceBytes, &response)
	if err != nil {
		return nil, errors.Join(errors.New("error in deserializing api model"), err)
	}
	return &response, nil
}

func (converter JsonMarshallingConverter[proto, db]) ApiToDb(source *proto) (*db, error) {
	var response db
	sourceBytes, err := json.Marshal(*source)
	if err != nil {
		return nil, errors.Join(errors.New("error in serializing api model"), err)
	}
	err = json.Unmarshal(sourceBytes, &response)
	if err != nil {
		return nil, errors.Join(errors.New("error in deserializing db model"), err)
	}
	return &response, nil
}

func (converter JsonMarshallingConverter[proto, db]) DbToApiSlice(source []*db) (_ []*proto, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error in marshalling DB collection to API collection: %v", r)
		}
	}()
	return lo.Map(source, func(item *db, i int) *proto {
		x, err := converter.DbToApi(item)
		if err != nil {
			panic(err)
		}
		return x
	}), nil
}

func (converter JsonMarshallingConverter[proto, db]) ApiToDbSlice(source []*proto) (_ []*db, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error in marshalling API collection to DB collection: %v", r)
		}
	}()
	return lo.Map(source, func(item *proto, i int) *db {
		x, err := converter.ApiToDb(item)
		if err != nil {
			panic(fmt.Sprintf("error in converting item at index %d: %v", i, err))
		}
		return x
	}), nil
}
