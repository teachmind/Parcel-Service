package carrier

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"parcel-service/internal/app/model"
	"testing"
	"time"
)

func TestRepository_UpdateCarrierRequest(t *testing.T) {
	sourceTime := time.Now()
	parcel := model.CarrierRequest{
		ParcelID: 1,
		CarrierID:    2,
		Status:  1,
	}

	statuses := model.Statuses{
		ParcelStatus: 2,
		AcceptStatus: 2,
		RejectStatus: 3,
	}

	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.RejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, statuses.ParcelStatus, sourceTime, parcel.ParcelID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectCommit()

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.True(t, errors.Is(err, nil))
	})

	t.Run("should return begin transaction failed", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin().WillReturnError(model.ErrTransaction)

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.True(t, errors.Is(err, model.ErrTransaction))
	})

	t.Run("UpdateAcceptQuery should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("UpdateAcceptQuery should return sql-error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.NotNil(t, err)
	})

	t.Run("updateRejectQuery should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.RejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("updateRejectQuery should return sql-error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.RejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.NotNil(t, err)
	})

	t.Run("updateParcelStatus should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.RejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, statuses.ParcelStatus, sourceTime, parcel.ParcelID).
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("updateParcelStatus should return sql-error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.RejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, statuses.ParcelStatus, sourceTime, parcel.ParcelID).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.NotNil(t, err)
	})

	t.Run("should return commit failed", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.AcceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(statuses.RejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, statuses.ParcelStatus, sourceTime, parcel.ParcelID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		m.ExpectCommit().WillReturnError(model.ErrTransaction)

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, statuses.AcceptStatus, statuses.RejectStatus, statuses.ParcelStatus, sourceTime)
		assert.True(t, errors.Is(err, model.ErrTransaction))
	})
}
