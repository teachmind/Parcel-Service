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
	insertParcelQuery  = `INSERT INTO parcel (user_id, carrier_id, source_address, destination_address, source_time, type, price, carrier_fee, company_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	getParcelListQuery = `SELECT * FROM parcel WHERE status=$1 LIMIT $2 OFFSET $3`
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
	// actualPhone := RMCodeAndSpace(user.PhoneNumber)
	fmt.Print(parcel.CarrierFee)
	if _, err := r.db.ExecContext(ctx, insertParcelQuery, parcel.UserID, 1, parcel.SourceAddress, parcel.DestinationAddress, parcel.SourceTime, parcel.ParcelType, 200.0, 180.0, 20.0); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}

func (r *repository) GetParcelsList(ctx context.Context, status int, limit int, offset int) ([]model.Parcel, error) {
	var parcels []model.Parcel
	//r.db.Select(&parcels, getParcelListQuery)
	if err := r.db.Select(&parcels, getParcelListQuery, status, limit, offset); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return nil, fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return nil, err
	}
	fmt.Println(parcels)
	return parcels, nil
}
