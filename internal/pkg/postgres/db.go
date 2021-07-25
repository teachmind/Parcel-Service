package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	defaultMacIleConn      = 10
	defaultMaxOpenConn     = 10
	defaultConnMaxLifetime = 30 * time.Minute
)

type Config struct {
	Host                     string
	Port                     string
	User                     string
	Password                 string
	Name                     string
	MaxIdleConn              int
	MaxOpenConn              int
	ConnMacLifeTimeInMinutes int
}

func New(conf *Config) (*sqlx.DB, error) {
	println(conf.Url())

	db, err := sqlx.Connect("postgres", conf.Url())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(conf.maxIdleConn())
	db.SetMaxOpenConns(conf.maxOpenConn())
	db.SetConnMaxLifetime(conf.connMaxLifeTime())

	return db, nil
}

func (c *Config) Url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name)
}

func (c *Config) maxIdleConn() int {
	if c.MaxIdleConn == 0 {
		return defaultMacIleConn
	}
	return c.MaxIdleConn
}

func (c *Config) maxOpenConn() int {
	if c.MaxOpenConn == 0 {
		return defaultMaxOpenConn
	}
	return c.MaxOpenConn
}

func (c *Config) connMaxLifeTime() time.Duration {
	if c.ConnMacLifeTimeInMinutes == 0 {
		return defaultConnMaxLifetime
	}
	return time.Duration(c.ConnMacLifeTimeInMinutes) * time.Minute
}
