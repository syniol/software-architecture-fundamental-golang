package storage

type Repository interface {
	SaveOne(entity Entity) error
}
