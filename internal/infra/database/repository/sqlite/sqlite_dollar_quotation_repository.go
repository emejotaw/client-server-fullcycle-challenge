package sqlite

import (
	"context"
	"time"

	"github.com/emejotaw/client-server-fullcycle-challenge/internal/infra/database/entity"
	"gorm.io/gorm"
)

type SqliteDollarQuotationRepository struct {
	db                *gorm.DB
	databaseTimeoutMs int
}

func NewSqliteRepository(db *gorm.DB, databaseTimeoutMs int) *SqliteDollarQuotationRepository {
	return &SqliteDollarQuotationRepository{db: db, databaseTimeoutMs: databaseTimeoutMs}
}

func (r *SqliteDollarQuotationRepository) Create(dollarQuotation *entity.DollarQuotation) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(r.databaseTimeoutMs))
	defer cancel()
	return r.db.WithContext(ctx).Create(dollarQuotation).Error
}
