package storage

import "os"

type Manager interface {
	SaveFile(name string, content []byte) error
}

func NewStorage() Manager {
	return newStorageUnix(os.Getenv("STORAGE_PATH"))
}
