package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database"
	"github.com/emejotaw/client-server-fullcycle-challenge/internal/service"
	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/client/httpclient"
	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/constants"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

const DATABASE_TIMEOUT_MS = 10
const HTTP_REQUEST_TIMEOUT_MS = 200

func NewServer(db *gorm.DB) *Server {

	os.Setenv(constants.DATABASE_TIMEOUT_MS_NAME, fmt.Sprintf("%d", DATABASE_TIMEOUT_MS))
	os.Setenv(constants.HTTP_REQUEST_TIMEOUT_MS_NAME, fmt.Sprintf("%d", HTTP_REQUEST_TIMEOUT_MS))

	return &Server{
		db: db,
	}
}

func main() {

	port := ":8080"
	db, err := database.GetConnection()

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(db)

	if err != nil {
		panic(err)
	}

	server := NewServer(db)

	log.Printf("Server will run at port %s", port)
	http.HandleFunc("/cotacao", server.GetDollarQuotation)
	http.ListenAndServe(port, nil)
}

func (s *Server) GetDollarQuotation(w http.ResponseWriter, r *http.Request) {

	client := httpclient.NewHttpClient()

	dollarQuotationDTO, err := client.GetDollarQuotation()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("server could not get dollar quotation, error: %v", err.Error())))
		return
	}

	service, err := service.NewDollarQuotationService(s.db)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error generating service, error: %v", err.Error())))
		return
	}

	err = service.Create(dollarQuotationDTO)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not create dollar quotation, error: %v", err.Error())))
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(dollarQuotationDTO)
}
