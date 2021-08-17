package carrier

import (
	"context"
	"errors"
	"parcel-service/internal/app/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
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
		assert.True(t, errors.Is(err, nil))
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
		m.ExpectExec("INSERT INTO carrier_requests (.+) VALUES (.+)").
			WithArgs().
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.InsertCarrierRequest(context.Background(), carrierRequest)
		assert.NotNil(t, err)
	})
}
