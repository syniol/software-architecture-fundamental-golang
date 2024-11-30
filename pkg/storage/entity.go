package storage

type Entity struct {
	ID            string
	FileName      string
	FileExtension string
	MetaData      map[string]interface{}
	Content       []byte
}
