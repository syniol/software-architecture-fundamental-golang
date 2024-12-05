package card

import (
	"encoding/json"
	"fmt"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/database"
)

type StudentRepo struct {
	dbClient *database.Database
}

func NewStudentCardRepository() database.RepositoryHandler[Card] {
	return StudentRepo{}
}

func (sr StudentRepo) CreateOne(card Card) error {
	cardJSON, err := json.Marshal(card)
	if err != nil {
		return err
	}

	sqlStmt := fmt.Sprintf(`INSERT INTO student.card (data) VALUES ('%s')`, string(cardJSON))

	_, err = sr.dbClient.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

func (sr StudentRepo) FindOneWithID(id int) database.Entity[Card] {
	return database.Entity[Card]{}
}

func (sr StudentRepo) UpdateOne(entity database.Entity[Card]) error {
	return nil
}

func (sr StudentRepo) DeleteOne(entity database.Entity[Card]) error {
	return nil
}
