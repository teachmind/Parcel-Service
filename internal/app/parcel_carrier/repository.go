package parcel_carrier

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"parcel-service/internal/app/model"
)
// sql query and error
const (
	errUniqueViolation = pq.ErrorCode("23505")
	//updateAcceptQuery = `UPDATE carrier_request SET status = $1 WHERE parcel_id = $2 AND carrier_id = $3`
	updateQuery = `UPDATE carrier_request 
							SET status = (CASE
											WHEN carrier_id = $4 THEN $1
											WHEN carrier_id != $4 THEN $2
										  END)
							WHERE parcel_id = $3`
)

type repository struct {
	db *sqlx.DB
}

// NewRepository initiates to assign carrier to parcel repository and return DB
func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest) error {
	if _, err := r.db.ExecContext(ctx, updateQuery, 1, 2, parcel.ParcelID, parcel.CarrierID); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}
