package storage

type RepositoryManager interface {
	SaveSingleFile(filename, extension string, content []byte) error
}
