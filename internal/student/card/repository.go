package card

import (
	"encoding/json"
	"fmt"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/database"
	"github.com/syniol/software-architecture-fundamental-golang/pkg/storage"
)

// StudentCardRepositoryManager is an example of multiple inheritance in Go
type StudentCardRepositoryManager interface {
	database.RepositoryManager[Card]
	storage.RepositoryManager

	// Add any custom method here and create an implementation on:
	// func (sr *StudentCardRepository) YourCustomFunc(param any): any
	findStudentCardWithNameAndID(name, studentID string) (*Card, error)
}

type StudentCardRepository struct {
	storageClient storage.Manager
	dbClient      *database.Database
}

func NewStudentCardRepository() (StudentCardRepositoryManager, error) {
	dbClient, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	return &StudentCardRepository{
		dbClient:      dbClient,
		storageClient: storage.NewStorage(),
	}, nil
}

func (sr *StudentCardRepository) SaveSingleFile(studentID, extension string, content []byte) error {
	err := sr.storageClient.SaveFile(fmt.Sprintf("%s.%s", studentID, extension), content)
	if err != nil {
		return err
	}

	return nil
}

func (sr *StudentCardRepository) CreateOne(card *Card) error {
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

// findStudentCardWithNameAndID is an example of additional custom repository methods specific to each domain
func (sr *StudentCardRepository) findStudentCardWithNameAndID(name, studentID string) (*Card, error) {
	return nil, nil
}
