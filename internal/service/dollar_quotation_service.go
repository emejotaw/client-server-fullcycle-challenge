package service

import (
	"log"
	"os"
	"strconv"

	"github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database/entity"
	"github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database/repository"
	"github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database/repository/sqlite"
	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/constants"
	"github.com/emejotaw/client-server-fullcycle-challenge/pkg/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DollarQuotationService struct {
	repository repository.DollarQuotationRepository
}

func NewDollarQuotationService(db *gorm.DB) (*DollarQuotationService, error) {

	databaseTimeoutMs, err := strconv.Atoi(os.Getenv(constants.DATABASE_TIMEOUT_MS_NAME))

	if err != nil {
		return nil, err
	}
	repository := sqlite.NewSqliteRepository(db, databaseTimeoutMs)
	return &DollarQuotationService{
		repository: repository,
	}, nil
}

func (s *DollarQuotationService) Create(dollarQuotationDTO *dto.DollarQuotationDTO) error {

	log.Printf("A new dollar quotation will be created with data: %v", dollarQuotationDTO)

	dollarQuotation := &entity.DollarQuotation{
		ID:         uuid.New().String(),
		Code:       dollarQuotationDTO.Code,
		Codein:     dollarQuotationDTO.Codein,
		Name:       dollarQuotationDTO.Name,
		High:       dollarQuotationDTO.High,
		Low:        dollarQuotationDTO.Low,
		VarBid:     dollarQuotationDTO.VarBid,
		PctChange:  dollarQuotationDTO.PctChange,
		Bid:        dollarQuotationDTO.Bid,
		Ask:        dollarQuotationDTO.Ask,
		Timestamp:  dollarQuotationDTO.Timestamp,
		CreateDate: dollarQuotationDTO.CreateDate,
	}

	return s.repository.Create(dollarQuotation)
}
