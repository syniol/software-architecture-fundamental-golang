package card

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/internal/student/card"
	"github.com/syniol/software-architecture-fundamental-golang/pkg/database"
)

func createStudentCardEndpointHandler(
	studentRepository database.RepositoryManager[card.Card],
	wr http.ResponseWriter,
	rq *http.Request,
) {
	reqBody, err := io.ReadAll(rq.Body)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		wr.Write([]byte(`{"error": "unexpected error occurred creating a card. Please try again or contact IT Support."}`))

		// Publishing error for internal logs
		log.Println(err.Error())

		return
	}

	var studentCardRequest card.Card
	err = json.Unmarshal(reqBody, &studentCardRequest)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		wr.Write([]byte(`{"error": "unexpected error occurred creating a card. Please try again or contact IT Support."}`))

		// Publishing error for internal logs
		log.Println(err.Error())

		return
	}

	studentCard, err := card.NewStudentCardWithDTO(reqBody)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		wr.Write([]byte(`{"error": "unexpected error occurred creating a card. Please try again or contact IT Support."}`))

		// Publishing error for internal logs
		log.Println(err.Error())

		return
	}

	err = studentRepository.CreateOne(
		card.NewStudentCard(
			studentCard.StudentID,
			studentCard.Name,
		),
	)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		wr.Write([]byte(`{"error": "unexpected error occurred creating a card. Please try again or contact IT Support."}`))

		// Publishing error for internal logs
		log.Println(err.Error())

		return
	}

	wr.WriteHeader(http.StatusAccepted)
	wr.Write([]byte(`{"status": "success"}`))
}
