package parcel

import (
	"context"
	"errors"
	"parcel-service/internal/app/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRepository_InsertParcel(t *testing.T) {
	parcel := model.Parcel{
		UserID:             1,
		SourceAddress:      "Dhaka Bangladesh",
		DestinationAddress: "Pabna Shadar",
		SourceTime:         time.Now(),
		ParcelType:         "Document",
		Price:              200.0,
		CarrierFee:         180.0,
		CompanyFee:         20.0,
	}

	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO parcel (.+) VALUES (.+)").
			WithArgs(parcel.UserID, parcel.SourceAddress, parcel.DestinationAddress, parcel.SourceTime, parcel.ParcelType, parcel.Price, parcel.CarrierFee, parcel.CompanyFee).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(sqlxDB)
		err := repo.InsertParcel(context.Background(), parcel)
		assert.Nil(t, err)
	})

	t.Run("should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO parcel (.+) VALUES (.+)").
			WithArgs(parcel.UserID, parcel.SourceAddress, parcel.DestinationAddress, parcel.SourceTime, parcel.ParcelType, parcel.Price, parcel.CarrierFee, parcel.CompanyFee).
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.InsertParcel(context.Background(), parcel)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return sql error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO parcel (.+) VALUES (.+)").
			WithArgs().
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.InsertParcel(context.Background(), parcel)
		assert.EqualError(t, err, "sql-error")
	})
}
