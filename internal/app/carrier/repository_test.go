package carrier

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"parcel-service/internal/app/model"
	"testing"
	"time"
)

func TestRepository_InsertCarrierRequest(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		carrierRequest := model.CarrierRequest{
			CarrierID: 1,
			ParcelID:  1,
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO carrier_request (.+) VALUES (.+)").
			WithArgs(carrierRequest.CarrierID, carrierRequest.ParcelID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(sqlxDB)
		err := repo.InsertCarrierRequest(context.Background(), carrierRequest)
		assert.Nil(t, err)
	})

	t.Run("should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		carrierRequest := model.CarrierRequest{
			CarrierID: 1,
			ParcelID:  1,
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO carrier_request (.+) VALUES (.+)").
			WithArgs(carrierRequest.CarrierID, carrierRequest.ParcelID).
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.InsertCarrierRequest(context.Background(), carrierRequest)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return sql error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		carrierRequest := model.CarrierRequest{
			CarrierID: 1,
			ParcelID:  1,
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO carrier_request (.+) VALUES (.+)").
			WithArgs().
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.InsertCarrierRequest(context.Background(), carrierRequest)
		assert.EqualError(t, err, "sql-error")
	})
}

func TestRepository_UpdateCarrierRequest(t *testing.T) {
	sourceTime := time.Now()
	parcel := model.CarrierRequest{
		ParcelID:  1,
		CarrierID: 2,
		Status:    1,
	}

	const acceptStatus, rejectStatus, parcelStatus int = 2, 3, 2

	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(acceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(rejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, parcelStatus, sourceTime, parcel.ParcelID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectCommit()

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
		assert.Nil(t, err)
	})

	t.Run("should return no rows error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("SELECT (.+) FROM parcel WHERE (.+)").
			WithArgs(parcel.ParcelID).
			WillReturnError(sql.ErrNoRows)
		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
		fmt.Println(err)
		assert.True(t, errors.Is(err, model.ErrNotFound))
	})

	t.Run("should return internal server error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin().WillReturnError(model.IntServerErr)

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
		assert.EqualError(t, err, "internal server error")
	})

	t.Run("UpdateAcceptQuery should return sql-error to update carrier_request table for accepting", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(acceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
		assert.EqualError(t, err, "sql-error")
	})


	t.Run("updateRejectQuery should return sql-error to update carrier_request table for rejecting", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(acceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(rejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
		assert.EqualError(t, err, "sql-error")
	})

	t.Run("updateParcelStatus should return sql-error to update parcel table with parcel status", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(acceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(rejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, parcelStatus, sourceTime, parcel.ParcelID).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
		assert.EqualError(t, err, "sql-error")
	})

	t.Run("should return commit failed", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectBegin()
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(acceptStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)").
			WithArgs(rejectStatus, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.CarrierID, parcelStatus, sourceTime, parcel.ParcelID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		m.ExpectCommit().WillReturnError(model.IntServerErr)

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
		assert.True(t, errors.Is(err, model.IntServerErr))
	})
}
