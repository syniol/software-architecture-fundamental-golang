package database

type RepositoryHandler[T any] interface {
	CreateOne(entity *T) error
	FindOneWithID(id int) Entity[T]
	UpdateOne(entity Entity[T]) error
	DeleteOne(entity Entity[T]) error
}
