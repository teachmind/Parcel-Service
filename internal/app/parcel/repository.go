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
	insertParcelQuery    = `INSERT INTO parcel (user_id, source_address, destination_address, source_time, type, price, carrier_fee, company_fee) VALUES (:user_id, :source_address, :destination_address, :source_time, :type, :price, :carrier_fee, :company_fee) RETURNING id, created_at, updated_at`
	fetchParcelByIDQuery = `SELECT * FROM parcel WHERE id = $1`
	updateParcelQuery    = `UPDATE parcel SET status = $1 WHERE id = $2`
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

func (r *repository) InsertParcel(ctx context.Context, parcel model.Parcel) (model.Parcel, error) {
	stmt, err := r.db.PrepareNamedContext(ctx, insertParcelQuery)

	if err != nil {
		log.Error().Err(err).Msgf("[InsertParcel] from statemtnt: %v", err)
		return model.Parcel{}, err
	}

	err = stmt.GetContext(ctx, &parcel, &parcel)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return model.Parcel{}, fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		log.Error().Err(err).Msgf("[FetchParcelByID] Error from getcontext 2: %v", err)
		return model.Parcel{}, err
	}

	return parcel, nil
}

func (r *repository) FetchParcelByID(ctx context.Context, parcelID int) (model.Parcel, error) {
	var parcel model.Parcel

	if err := r.db.GetContext(ctx, &parcel, fetchParcelByIDQuery, parcelID); err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msgf("[FetchParcelByID] failed to fetch parcel Error: %v", err)
			return model.Parcel{}, fmt.Errorf("parcel with the ID %d is not found. :%w", parcelID, model.ErrNotFound)
		}

		return model.Parcel{}, err
	}

	return parcel, nil
}

func (r *repository) UpdateParcel(ctx context.Context, parcel model.Parcel) error {
	result, err := r.db.ExecContext(ctx, updateParcelQuery, parcel.Status, parcel.ID)

	if err != nil {
		log.Error().Err(err).Msgf("[UpdateParcel] failed to update parcel Error: %v", err)

		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}

		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("%v :%w", err, model.ErrInvalid)
	}

	if rows == 0 {
		return fmt.Errorf("parcel %d not updated, please provide valid ID. :%w", parcel.ID, model.ErrNotFound)
	}

	return nil
}
