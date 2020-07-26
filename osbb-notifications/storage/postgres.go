package storage

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

const (
	pgStr        = "host=%v port=%v user=%v password=%v dbname=%v sslmode=disable"
	dbDriverName = "postgres"
)

// Config of Postgres DB
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// ConnectToDB() makes connect to Postgres DB
func ConnectToDB(pg *DBConfig) (*sql.DB, error) {
	pgConfig := getDBConfigString(pg)

	database, err := sql.Open(dbDriverName, pgConfig)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}

// getDBConfigString() creates connStr for sql.Open
func getDBConfigString(pg *DBConfig) string {
	return fmt.Sprintf(pgStr, pg.Host, pg.Port, pg.User, pg.Password, pg.DBName)
}

const queryInsert = `INSERT INTO public.notifications(id, message) VALUES ($1,$2)`

//Notifications{} contains pointer to database
type Notifications struct {
	db *sql.DB
}

// NewNotifications() returns Notifications with db
func NewNotifications(data *sql.DB) *Notifications {
	return &Notifications{db: data}
}

// Add() adds new user to database
func (n *Notifications) Add(message string) error {
	ui, err := uuid.NewRandom()

	if err != nil {
		return fmt.Errorf("uuid.NewRandom error: %w", err)
	}

	_, err = n.db.Exec(queryInsert, ui, message)

	return err
}
