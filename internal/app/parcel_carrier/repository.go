package parcel_carrier

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"parcel-service/internal/app/model"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)
// sql query and error
const (
	errUniqueViolation = pq.ErrorCode("23505")
	updateParcelStatus = `UPDATE parcel SET carrier_id = $1, status = $2, source_time = $3 WHERE id = $4`
	updateAcceptQuery = `UPDATE carrier_request SET status = $1 WHERE parcel_id = $2 AND carrier_id = $3`
	updateRejectQuery = `UPDATE carrier_request SET status = $1 WHERE parcel_id = $2 AND carrier_id != $3`
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

func (r *repository) UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest, acceptStatus int, rejectStatus int, parcelStatus int, sourceTime time.Time) error {
	//starting db transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCarrierStatus] has failed to begin transactions.")
		return err
	}
	//accept status update for carrier request table
	if _, err = tx.ExecContext(ctx, updateAcceptQuery, acceptStatus, parcel.ParcelID, parcel.CarrierID); err != nil {
		tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			log.Error().Err(err).Msg("[UpdateCarrierRequest] has failed to update status of carrier_rquest table for accept.")
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	if _, err := tx.ExecContext(ctx, updateRejectQuery, rejectStatus, parcel.ParcelID, parcel.CarrierID); err != nil {
		tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			log.Error().Err(err).Msg("[UpdateCarrierRequest] has failed to update reject status for carrier_request table.")
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	if _, err := tx.ExecContext(ctx, updateParcelStatus, parcel.CarrierID, parcelStatus, sourceTime, parcel.ParcelID); err != nil {
		tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == errUniqueViolation {
			return fmt.Errorf("%v :%w", err, model.ErrInvalid)
		}
		return err
	}
	if err=tx.Commit();err!=nil{
		tx.Rollback()
		log.Error().Err(err).Msg("[UpdateCarrierRequest] failed to commit transaction")
		return err
	}
	return nil
}
