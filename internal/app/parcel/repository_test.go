package parcel

import (
	"context"
	"database/sql"
	"errors"
	"parcel-service/internal/app/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRepository_InsertUser(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			PhoneNumber: "01738799349",
			Password:    "123456",
			CategoryID:  1,
		}
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO users (.+) VALUES (.+)").
			WithArgs(user.PhoneNumber, user.Password, user.CategoryID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(sqlxDB)
		err := repo.InsertUser(context.Background(), user)
		assert.True(t, errors.Is(err, nil))
	})

	t.Run("should return unique key violation error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			PhoneNumber: "01738799349",
			Password:    "123456",
			CategoryID:  1,
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO users (.+) VALUES (.+)").
			WithArgs(user.PhoneNumber, user.Password, user.CategoryID).
			WillReturnError(&pq.Error{Code: "23505"})

		repo := NewRepository(sqlxDB)
		err := repo.InsertUser(context.Background(), user)
		assert.True(t, errors.Is(err, model.ErrInvalid))
	})

	t.Run("should return sql error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			PhoneNumber: "01738799349",
			Password:    "123456",
			CategoryID:  1,
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectExec("INSERT INTO users (.+) VALUES (.+)").
			WithArgs(user.PhoneNumber, user.Password).
			WillReturnError(errors.New("sql-error"))

		repo := NewRepository(sqlxDB)
		err := repo.InsertUser(context.Background(), user)
		assert.NotNil(t, err)
	})
}

func TestRepository_GetUserByPhone(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		user := model.User{
			ID:          1,
			PhoneNumber: "01738799349",
			Password:    "123456",
			CategoryID:  1,
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("01738799349").
			WillReturnRows(sqlmock.NewRows([]string{"id", "phone_number", "password", "category_id"}).
				AddRow(1, "01738799349", "123456", 1))

		repo := NewRepository(sqlxDB)
		result, err := repo.GetUserByPhone(context.Background(), "01738799349")

		assert.Nil(t, err)
		assert.EqualValues(t, user, result)
	})

	t.Run("should return no rows error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone_number").
			WillReturnError(sql.ErrNoRows)
		repo := NewRepository(sqlxDB)
		_, err := repo.GetUserByPhone(context.Background(), "phone_number")
		assert.True(t, errors.Is(err, model.ErrNotFound))
	})

	t.Run("should return error", func(t *testing.T) {
		db, m, _ := sqlmock.New()
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		m.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)").
			WithArgs("phone_number").
			WillReturnError(errors.New("sql-error"))
		repo := NewRepository(sqlxDB)
		_, err := repo.GetUserByPhone(context.Background(), "phone_number")
		assert.NotNil(t, err)
	})
}
