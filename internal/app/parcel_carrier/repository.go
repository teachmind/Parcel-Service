package parcel_carrier

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
	updateParcelStatus = `UPDATE parcel SET carrier_id = $1, status = $2 WHERE id = $3`
	updateQuery = `UPDATE carrier_request SET status = (CASE WHEN carrier_id = $2 THEN $3 WHEN carrier_id != $2 THEN $4 END) WHERE parcel_id = $1`
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

func (r *repository) UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest, status model.ParcelStatus) error {

	if _, err := r.db.ExecContext(ctx, updateQuery, parcel.ParcelID, parcel.CarrierID, status.Accept, status.Reject); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		if _, err := r.db.ExecContext(ctx, updateParcelStatus, parcel.CarrierID, status.ParcelStatus, parcel.ParcelID); err != nil {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}
