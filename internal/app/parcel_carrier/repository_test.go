package parcel_carrier

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
	var sourceTime = time.Now()
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		parcel := model.CarrierRequest{
			ParcelID: 1,
			CarrierID:    2,
			Status:  1,
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(2, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(3, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, 2, sourceTime, parcel.ParcelID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectCommit()

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, 2, 3, 2, sourceTime)
		assert.True(t, errors.Is(err, nil))
	})

	t.Run("should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		parcel := model.CarrierRequest{
			ParcelID: 1,
			CarrierID:    2,
			Status:  1,
		}
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request").
			WillReturnError(&pq.Error{Code: "23505"})
		m.ExpectExec("UPDATE parcel").
			WillReturnError(&pq.Error{Code: "23505"})
		m.ExpectRollback()

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, 2, 3, 2, time.Now())
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return sql-error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		parcel := model.CarrierRequest{
			ParcelID: 1,
			CarrierID:    2,
			Status:  1,
		}
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request").
			WillReturnError(errors.New("sql-error"))
		m.ExpectExec("UPDATE parcel").
			WillReturnError(errors.New("sql-error"))
		m.ExpectRollback()

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, 2, 3, 2, time.Now())
		assert.NotNil(t, err)
	})
}
