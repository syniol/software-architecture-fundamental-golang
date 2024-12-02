package database

type RepositoryHandler[T any] interface {
	SaveOne(entity Entity[T]) error
}
