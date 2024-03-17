package models

type Object struct {
	Filename     string
	FileLocation string
	ContentType  string
	ContentSize  int // int32 4,294,967,295 or int64 9,223,372,036,854,775,807
}
