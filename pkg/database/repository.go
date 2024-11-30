package database

type Repository[T any] interface {
	SaveOne(entity Entity[T]) error
}
