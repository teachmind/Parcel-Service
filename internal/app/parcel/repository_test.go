package parcel

import (
	"context"
	"database/sql"
	"errors"
	"parcel-service/internal/app/model"
	"regexp"
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
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
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
		assert.NotNil(t, err)
	})
}

func TestRepository_GetParcelsList(t *testing.T) {

	var status int
	var limit int
	var offset int

	parcels := []model.Parcel{
		{
			ID:                 0,
			UserID:             1,
			CarrierID:          0,
			Status:             1,
			SourceAddress:      "Dhaka Bangladesh",
			DestinationAddress: "Pabna Shadar",
			ParcelType:         "Document",
			Price:              200,
			CarrierFee:         180,
			CompanyFee:         20,
		}, {
			ID:                 0,
			UserID:             2,
			CarrierID:          0,
			Status:             1,
			SourceAddress:      "Dhaka Bangladesh",
			DestinationAddress: "Pabna Shadar",
			ParcelType:         "Document",
			Price:              200,
			CarrierFee:         180,
			CompanyFee:         20,
		}}

	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")

		status = 1
		limit = 2
		offset = 0

		m.ExpectQuery(regexp.QuoteMeta("SELECT user_id, status, source_address, destination_address, type, price, carrier_fee, company_fee FROM parcel WHERE status=$1 LIMIT $2 OFFSET $3")).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "status", "source_address", "destination_address", "type", "price", "carrier_fee", "company_fee"}).
				AddRow(parcels[0].UserID, parcels[0].Status, parcels[0].SourceAddress, parcels[0].DestinationAddress, parcels[0].ParcelType, parcels[0].Price, parcels[0].CarrierFee, parcels[0].CompanyFee).
				AddRow(parcels[1].UserID, parcels[1].Status, parcels[1].SourceAddress, parcels[1].DestinationAddress, parcels[1].ParcelType, parcels[1].Price, parcels[1].CarrierFee, parcels[1].CompanyFee))

		repo := NewRepository(sqlxDB)
		result, err := repo.GetParcelsList(context.Background(), status, limit, offset)

		assert.Nil(t, err)
		assert.EqualValues(t, parcels, result)
	})

	t.Run("should return success with 0 rows", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")

		status = 0
		limit = 0
		offset = 0

		m.ExpectQuery("^SELECT (.+) FROM parcel WHERE (.+)").
			WithArgs(0, 0, 0).
			WillReturnRows(sqlmock.NewRows([]string{}))

		repo := NewRepository(sqlxDB)
		res, err := repo.GetParcelsList(context.Background(), status, limit, offset)

		assert.Empty(t, res)
		assert.Nil(t, err)
	})

	t.Run("should return no rows error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")

		status = 0
		limit = 0
		offset = -1

		m.ExpectQuery("^SELECT (.+) FROM parcel WHERE (.+)").
			WithArgs(0, 0, -1).
			WillReturnError(sql.ErrNoRows)

		repo := NewRepository(sqlxDB)
		_, err := repo.GetParcelsList(context.Background(), status, limit, offset)

		assert.EqualError(t, err, "parcel list for offset -1 is not found. :"+sql.ErrNoRows.Error())
	})

	t.Run("should return sql error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")

		m.ExpectQuery("^SELECT (.+) FROM parcel WHERE (.+)").
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		_, err := repo.GetParcelsList(context.Background(), status, limit, offset)

		assert.EqualError(t, err, "sql-error")
	})

}

func TestRepository_FetchParcelByID(t *testing.T) {
	parcel := model.Parcel{
		ID:                 1,
		UserID:             1,
		SourceAddress:      "Dhaka Bangladesh",
		DestinationAddress: "Pabna Shadar",
		SourceTime:         time.Now(),
		ParcelType:         "Document",
		Price:              200.0,
		CarrierFee:         180.0,
		CompanyFee:         20.0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM parcel WHERE (.+)").
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "source_address", "destination_address", "source_time", "type", "price", "carrier_fee", "company_fee", "created_at", "updated_at"}).
				AddRow(parcel.ID, parcel.UserID, parcel.SourceAddress, parcel.DestinationAddress, parcel.SourceTime, parcel.ParcelType, parcel.Price, parcel.CarrierFee, parcel.CompanyFee, parcel.CreatedAt, parcel.UpdatedAt))

		repo := NewRepository(sqlxDB)
		result, err := repo.FetchParcelByID(context.Background(), parcel.ID)

		assert.Nil(t, err)
		assert.EqualValues(t, parcel, result)
	})

	t.Run("should return no rows error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM parcel WHERE (.+)").
			WithArgs(1).
			WillReturnError(sql.ErrNoRows)
		repo := NewRepository(sqlxDB)
		_, err := repo.FetchParcelByID(context.Background(), parcel.ID)
		assert.True(t, errors.Is(err, model.ErrNotFound))
	})

	t.Run("should return error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM parcel WHERE (.+)").
			WithArgs(1).
			WillReturnError(errors.New("sql-error"))
		repo := NewRepository(sqlxDB)
		_, err := repo.FetchParcelByID(context.Background(), parcel.ID)
		assert.EqualError(t, err, "sql-error")
	})
}

func TestRepository_UpdateParcel(t *testing.T) {
	parcel := model.Parcel{
		ID:                 1,
		UserID:             1,
		SourceAddress:      "Dhaka Bangladesh",
		DestinationAddress: "Pabna Shadar",
		SourceTime:         time.Now(),
		ParcelType:         "Document",
		Price:              200.0,
		CarrierFee:         180.0,
		CompanyFee:         20.0,
		Status:             2,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.Status, parcel.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateParcel(context.Background(), parcel)
		assert.Nil(t, err)
	})

	t.Run("should return invalid ID", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.Status, parcel.ID).
			WillReturnResult(sqlmock.NewResult(1, 0))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateParcel(context.Background(), parcel)
		assert.True(t, errors.Is(err, model.ErrNotFound))
	})

	t.Run("should return rows error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs(parcel.Status, parcel.ID).
			WillReturnResult(sqlmock.NewErrorResult(model.ErrInvalid))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateParcel(context.Background(), parcel)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs().
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.UpdateParcel(context.Background(), parcel)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return sql error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("UPDATE parcel SET (.+) WHERE (.+)").
			WithArgs().
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.UpdateParcel(context.Background(), parcel)
		assert.EqualError(t, err, "sql-error")
	})
}
