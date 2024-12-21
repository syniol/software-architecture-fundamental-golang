package card

import (
	"encoding/json"
	"time"

	"github.com/syniol/software-architecture-fundamental-golang/pkg"
)

type Card struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	IssueDate string `json:"issueDate"`
}

type DTO struct {
	StudentID string `json:"student_id"`
	Name      string `json:"full_name"`
}

func NewStudentCard(studentID, name string) *Card {
	return &Card{
		StudentID: studentID,
		Name:      name,
		IssueDate: pkg.NewDateTime().Format(time.DateOnly),
	}
}

func NewStudentCardWithDTO(request []byte) (*Card, error) {
	var studentCardDTO DTO
	err := json.Unmarshal(request, &studentCardDTO)
	if err != nil {
		return nil, err
	}

	return NewStudentCard(
		studentCardDTO.StudentID,
		studentCardDTO.Name,
	), nil
}
