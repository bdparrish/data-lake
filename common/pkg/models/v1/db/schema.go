package db

import "github.com/google/uuid"

type Object struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key;unique;default:gen_random_uuid()" json:"id"`
	FileName     string    `gorm:"type:varchar(255);not null" json:"file_name"`
	FileLocation string    `gorm:"type:varchar(255);not null" json:"file_location"`
	ContentType  string    `gorm:"type:varchar(255);not null" json:"content_type"`
	ContentSize  int32     `gorm:"type:int" json:"content_size"`
}
