package carrier

import (
	"context"
	"fmt"
	"parcel-service/internal/app/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// sql query and error
const (
	errUniqueViolation = pq.ErrorCode("23505")
	insertCarrierQuery = `INSERT INTO carrier_request (carrier_id, parcel_id) VALUES ($1, $2)`
)

type repository struct {
	db *sqlx.DB
}

// NewRepository initiates to perform carrier actions
func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertCarrierRequest(ctx context.Context, request model.CarrierRequest) error {
	if _, err := r.db.ExecContext(ctx, insertCarrierQuery, request.CarrierID, request.ParcelID); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}
