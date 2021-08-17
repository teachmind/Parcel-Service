package parcel_carrier

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"parcel-service/internal/app/model"
	"testing"
	"time"
)

func TestRepository_UpdateCarrierRequest(t *testing.T) {
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
		m.ExpectExec("update carrier_request set (.+) where (.+) and (.+)")

		m.ExpectRollback()

		m.ExpectQuery("UPDATE carrier_request SET (.+) WHERE (.+) AND (.+)")
			WithArgs(2, parcel.ParcelID, parcel.CarrierID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateCarrierRequest(context.Background(), parcel, 2, 3, 2, time.Now())
		assert.True(t, errors.Is(err, nil))
	})
}
