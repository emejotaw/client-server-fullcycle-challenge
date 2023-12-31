package repository

import "github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database/entity"

type DollarQuotationRepository interface {
	Create(*entity.DollarQuotation, int) error
}
