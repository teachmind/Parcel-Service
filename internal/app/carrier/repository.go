package carrier

import (
	"context"
	"fmt"
	"parcel-service/internal/app/model"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// sql query and error
const (
	errUniqueViolation = pq.ErrorCode("23505")
	updateAcceptQuery  = `UPDATE carrier_request SET status = $1 WHERE parcel_id = $2 AND carrier_id = $3`
	updateRejectQuery  = `UPDATE carrier_request SET status = $1 WHERE parcel_id = $2 AND carrier_id != $3`
	updateParcelStatus = `UPDATE parcel SET carrier_id = $1, status = $2, source_time = $3 WHERE id = $4`
	insertCarrierQuery = `INSERT INTO carrier_request (carrier_id, parcel_id) VALUES ($1, $2)`
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

func (r *repository) InsertCarrierRequest(ctx context.Context, request model.CarrierRequest) error {
	if _, err := r.db.ExecContext(ctx, insertCarrierQuery, request.CarrierID, request.ParcelID); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	return nil
}

func (r *repository) UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest, acceptStatus int, rejectStatus int, parcelStatus int, sourceTime time.Time) error {
	//starting db transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCarrierStatus] Internal Server Error.")
		return fmt.Errorf("%v", err)
	}
	//accept status update for carrier request table
	if _, err = tx.ExecContext(ctx, updateAcceptQuery, acceptStatus, parcel.ParcelID, parcel.CarrierID); err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[UpdateCarrierStatus] failed to update carrier_request table to accept: %v", err)
		return err
	}
	if _, err := tx.ExecContext(ctx, updateRejectQuery, rejectStatus, parcel.ParcelID, parcel.CarrierID); err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[UpdateCarrierStatus] failed to update carrier_request table to reject: %v", err)
		return err
	}
	if _, err := tx.ExecContext(ctx, updateParcelStatus, parcel.CarrierID, parcelStatus, sourceTime, parcel.ParcelID); err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[UpdateCarrierStatus] failed to update parcel table to update status: %v", err)
		return err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("[UpdateCarrierRequest] Failed to commit")
		return fmt.Errorf("%v :%w", err, model.IntServerErr)
	}
	return nil
}
