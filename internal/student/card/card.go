package card

import (
	"time"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/lib"
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

func NewStudentCard(studentID, name string) Card {
	return Card{
		StudentID: studentID,
		Name:      name,
		IssueDate: lib.NewDateTime().Format(time.DateOnly),
	}
}
