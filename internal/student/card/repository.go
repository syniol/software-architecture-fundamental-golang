package card

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/database"
	"github.com/syniol/software-architecture-fundamental-golang/pkg/storage"
)

type StudentRepository struct {
	DatabaseRepositoryManager database.RepositoryManager[Card]
	StorageRepositoryManager  storage.RepositoryManager

	storageClient storage.Manager
	dbClient      *database.Database
}

func NewStudentCardRepository() (*StudentRepository, error) {
	dbClient, err := database.NewDatabase(context.Background())
	if err != nil {
		return nil, err
	}

	return &StudentRepository{
		dbClient:      dbClient,
		storageClient: storage.NewStorage(),
	}, nil
}

func (sr *StudentRepository) SaveSingleFile(studentID, extension string, content []byte) error {
	err := sr.storageClient.SaveFile(fmt.Sprintf("%s.%s", studentID, extension), content)
	if err != nil {
		return err
	}

	return nil
}

func (sr *StudentRepository) CreateOne(card *Card) error {
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
