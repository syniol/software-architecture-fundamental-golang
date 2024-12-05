package card

import (
	"encoding/json"
)

type Card struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Photo     []byte `json:"photo"`
}

type DTO struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Photo     []byte `json:"photo"`
}

func NewStudentCard(studentID, name string, photo []byte) Card {
	return Card{
		StudentID: studentID,
		Name:      name,
		Photo:     photo,
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
