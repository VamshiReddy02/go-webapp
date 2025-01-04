package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func OpenDB(dns string) (*sql.DB, error ) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) ConnectToDB() (*sql.DB, error)  {
	connection, err := OpenDB(app.dsn)
	if err != nil {
		return nil, err
	}

	log.Println("connected to postgres")
	return connection, nil
}