package card

import (
	"bytes"
	"log"
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/internal/student/card"
	"github.com/syniol/software-architecture-fundamental-golang/pkg"
)

func NewRESTfulCreateStudentCardEndpoint() (path string, handler pkg.EndpointHandler) {
	return "/v1/student/card", pkg.NewRESTfulEndpointAuthenticatedHandler(
		func(wr http.ResponseWriter, rq *http.Request) {
			studentRepository, err := card.NewStudentCardRepository()
			if err != nil {
				wr.WriteHeader(http.StatusInternalServerError)
				wr.Write([]byte(`{ "error": "error establishing connection to database" }`))

				return
			}

			createStudentCardEndpointHandler(studentRepository, wr, rq)
		},
		http.MethodPost,
	)
}

func NewRESTfulStudentCardPhotoUploadEndpoint() (path string, handler pkg.EndpointHandler) {
	return "/v1/student/card/photo", pkg.NewRESTfulEndpointAuthenticatedHandler(
		func(wr http.ResponseWriter, rq *http.Request) {
			var imgBuffer bytes.Buffer
			if _, err := imgBuffer.ReadFrom(rq.Body); err != nil {
				wr.WriteHeader(http.StatusBadRequest)
				log.Println(err.Error())

				return
			}

			if err := rq.Body.Close(); err != nil {
				wr.WriteHeader(http.StatusInternalServerError)
				log.Println(err.Error())

				return
			}
			studentCardRepository, err := card.NewStudentCardRepository()
			if err != nil {
				wr.WriteHeader(http.StatusInternalServerError)
				log.Println(err.Error())

				return
			}

			if err = studentCardRepository.SaveSingleFile(
				"myidhere",
				"jpg",
				[]byte(imgBuffer.String()),
			); err != nil {
				wr.WriteHeader(http.StatusInternalServerError)
				log.Println(err.Error())

				return
			}

			wr.WriteHeader(http.StatusAccepted)
			wr.Write([]byte(`{ "message": "Photo is uploaded successfully" }`))
		},
		http.MethodPut,
	)
}
