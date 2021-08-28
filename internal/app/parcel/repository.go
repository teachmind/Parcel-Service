package parcel

import (
	"context"
	"database/sql"
	"fmt"
	"parcel-service/internal/app/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// SQL Query and error
const (
	errUniqueViolation   = pq.ErrorCode("23505")
	getParcelListQuery   = `SELECT * FROM parcel WHERE status=$1 LIMIT $2 OFFSET $3`
	insertParcelQuery    = `INSERT INTO parcel (user_id, source_address, destination_address, source_time, type, price, carrier_fee, company_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	fetchParcelByIDQuery = `SELECT * FROM parcel WHERE id = $1`
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

func (r *repository) FetchParcelByID(ctx context.Context, parcelID int) (model.Parcel, error) {
	var parcel model.Parcel

	if err := r.db.GetContext(ctx, &parcel, fetchParcelByIDQuery, parcelID); err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msgf("[FetchParcelByID] failed to fetch parcel Error: %v", err)
			return model.Parcel{}, fmt.Errorf("parcel %d is not found. :%w", parcelID, model.ErrNotFound)
		}

		return model.Parcel{}, err
	}

	return parcel, nil
}
