package parcel

import (
	"context"
	"fmt"
	"parcel-service/internal/app/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// SQL Query and error
const (
	errUniqueViolation = pq.ErrorCode("23505")
	insertParcelQuery  = `INSERT INTO parcel (user_id, source_address, destination_address, source_time, type, price, carrier_fee, company_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
)

type repository struct {
	db *sqlx.DB
}

// NewRepository initiates parcel repository and returns DB
func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertParcel(ctx context.Context, parcel model.Parcel) error {
	if _, err := r.db.ExecContext(ctx,
		insertParcelQuery,
		parcel.UserID,
		parcel.SourceAddress,
		parcel.DestinationAddress,
		parcel.SourceTime,
		parcel.ParcelType,
		parcel.Price,
		parcel.CarrierFee,
		parcel.CompanyFee); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}
