package sqlite

import (
	"context"
	"time"

	"github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database/entity"
	"gorm.io/gorm"
)

type SqliteDollarQuotationRepository struct {
	db *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) *SqliteDollarQuotationRepository {
	return &SqliteDollarQuotationRepository{db: db}
}

func (r *SqliteDollarQuotationRepository) Create(dollarQuotation *entity.DollarQuotation, databaseTimeoutMs int) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(databaseTimeoutMs))
	defer cancel()
	return r.db.WithContext(ctx).Create(dollarQuotation).Error
}
