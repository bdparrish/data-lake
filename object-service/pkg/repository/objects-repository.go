package repository

import (
	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/log"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ObjectRepository is the interface to handle CRUD operations for object model.
type ObjectRepository[T db.Object] interface {
	GetAll(page int, size int) ([]*T, error)
	Get(uuid uuid.UUID) (*T, error)
	Insert(record *T) (*T, error)
	Update(record *T) (*T, error)
	Delete(uuid uuid.UUID) (*T, error)
}

// ObjectRepositoryImpl is the struct to support the ObjectRepository methods.
type ObjectRepositoryImpl struct {
	Logger log.Logger
	Conf   *config.Config
	DB     *gorm.DB
}

// NewObjectRepository constructor for a new object repository
func NewObjectRepository(logger log.Logger, conf *config.Config, store *gorm.DB) ObjectRepositoryImpl {
	return ObjectRepositoryImpl{
		Logger: logger,
		Conf:   conf,
		DB:     store,
	}
}

func (repo ObjectRepositoryImpl) GetAll(page int, size int) ([]*db.Object, error) {
	offset := (page - 1) * size
	var results []*db.Object
	if err := repo.DB.
		Model(&db.Object{}).
		Limit(size).
		Offset(offset).
		Find(&results).
		Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo ObjectRepositoryImpl) Get(id uuid.UUID) (*db.Object, error) {
	var result db.Object
	if err := repo.DB.
		Model(&db.Object{}).
		Where("id = ?", id).
		First(&result).
		Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo ObjectRepositoryImpl) Insert(record *db.Object) (*db.Object, error) {
	if err := repo.DB.
		Model(&db.Object{}).
		Create(record).
		Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (repo ObjectRepositoryImpl) Update(record *db.Object) (*db.Object, error) {
	if err := repo.DB.
		Model(&db.Object{}).
		Where("id = ?", record.Id).
		Updates(record).
		Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (repo ObjectRepositoryImpl) Delete(id uuid.UUID) (*db.Object, error) {
	var result db.Object
	if err := repo.DB.
		Model(&db.Object{}).
		Where("id = ?", id).
		Delete(&result).
		Error; err != nil {
		return nil, err
	}
	return &result, nil
}
