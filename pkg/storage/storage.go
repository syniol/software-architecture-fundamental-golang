package storage

import "os"

type Partitioner interface {
	SaveFile(name string, content []byte) error
}

func NewStorage() Partitioner {
	return newStorageUnix(os.Getenv("STORAGE_PATH"))
}
