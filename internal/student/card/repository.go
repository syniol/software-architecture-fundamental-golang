package card

import "github.com/syniol/software-architecture-fundamental-golang/pkg/database"

var CardEntity database.Entity[Card]

type StudentRepo struct {
	*database.Database
}

func (sr StudentRepo) SaveOne(entity database.Entity[Card]) error {
	return nil
}

func NewStudentCardRepository() database.RepositoryHandler[Card] {
	return StudentRepo{}
}
