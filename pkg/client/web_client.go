package client

import "github.com/emejotaw/client-server-fullcycle-challenge/pkg/dto"

type WebClient interface {
	GetDollarQuotation() (*dto.DollarQuotationDTO, error)
}
