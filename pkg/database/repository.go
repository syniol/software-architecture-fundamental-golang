package database

type RepositoryManager[T any] interface {
	CreateOne(entity *T) error
}
