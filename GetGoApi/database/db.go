package database

import (
	"GetGoApi/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type PostDB interface {
	Open() error
	Close() error
	GetOrders() ([]*model.Order, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {

	pg, err := sqlx.Open("postgres", "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris")

	if err != nil {
		return err

	}
	log.Println("Connected to Database")

	d.db = pg
	return nil

}

func (d *DB) Close() error {

	return d.db.Close()

}

type OrdersDB interface {
	Open() error
	Close() error

	GetOrders() ([]*model.Order, error)
}
