package card

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID               string
	StudentID        string
	Name             string
	Photo            []byte
	CreatedTimestamp time.Time
}

type DTO struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Photo     []byte `json:"photo"`
}

func NewStudentCard(studentID, name string, photo []byte) Card {
	return Card{
		ID:               uuid.New().String(),
		StudentID:        studentID,
		Name:             name,
		Photo:            photo,
		CreatedTimestamp: time.Now(),
	}
}

func NewStudentCardWithCardDTO(dto *DTO) Card {
	return NewStudentCard(
		dto.StudentID,
		dto.Name,
		dto.Photo,
	)
}

func NewStudentCardDTO(input []byte) (error, *DTO) {
	var cardDTO DTO

	err := json.Unmarshal(input, &cardDTO)
	if err != nil {
		return err, nil
	}

	return nil, &cardDTO
}
