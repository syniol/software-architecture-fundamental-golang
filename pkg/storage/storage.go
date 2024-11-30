package storage

type Partitioner interface {
	SaveFile(name, path string, content []byte) error
}

func NewStorage() Partitioner {
	return NewStorageUnix()
}
