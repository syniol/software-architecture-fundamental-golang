package card

import (
	"bytes"
	"github.com/syniol/software-architecture-fundamental-golang/pkg/restful"
	"log"
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/internal/student/card"
)

func NewRESTfulCreateStudentCardEndpoint() (path string, handler restful.EndpointHandler) {
	return "/v1/student/card", restful.NewRESTfulEndpointAuthenticatedHandler(
		func(wr http.ResponseWriter, rq *http.Request) {
			studentRepository, err := card.NewStudentCardRepository()
			if err != nil {
				wr.WriteHeader(http.StatusInternalServerError)
				wr.Write([]byte(`{ "error": "error establishing a connection to database" }`))

				return
			}

			createStudentCardEndpointHandler(studentRepository, wr, rq)
		},
		http.MethodPost,
	)
}

func NewRESTfulStudentCardPhotoUploadEndpoint() (path string, handler restful.EndpointHandler) {
	return "/v1/student/card/photo", restful.NewRESTfulEndpointAuthenticatedHandler(
		func(wr http.ResponseWriter, rq *http.Request) {
			var imgBuffer bytes.Buffer
			if _, err := imgBuffer.ReadFrom(rq.Body); err != nil {
				wr.WriteHeader(http.StatusBadRequest)
				wr.Write([]byte(`{ "message": "There is an issue with submitted photograph, please contact the administrator" }`))

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
				imgBuffer.Bytes(),
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
