package card

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/database"
	"github.com/syniol/software-architecture-fundamental-golang/pkg/storage"
)

type StudentRepo struct {
	database.RepositoryHandler[Card]

	dbClient      *database.Database
	storageClient storage.Partitioner
}

func NewStudentCardRepository() (*StudentRepo, error) {
	dbClient, err := database.NewDatabase(context.Background())
	if err != nil {
		return nil, err
	}

	return &StudentRepo{
		dbClient:      dbClient,
		storageClient: storage.NewStorage(),
	}, nil
}

func (sr *StudentRepo) UploadPhotoID(studentID string, content []byte) error {
	err := sr.storageClient.SaveFile(fmt.Sprintf("%s.jpg", studentID), content)
	if err != nil {
		return err
	}

	return nil
}

func (sr *StudentRepo) CreateOne(card *Card) error {
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

func (sr *StudentRepo) FindOneWithID(id int) database.Entity[Card] {
	return database.Entity[Card]{}
}

func (sr *StudentRepo) FindOneWithStudentID(studentID string) database.Entity[Card] {
	return database.Entity[Card]{}
}

func (sr *StudentRepo) UpdateOne(entity database.Entity[Card]) error {
	return nil
}

func (sr *StudentRepo) DeleteOne(entity database.Entity[Card]) error {
	return nil
}
