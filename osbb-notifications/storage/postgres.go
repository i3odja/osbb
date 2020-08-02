package storage

import (
	"database/sql"
	"fmt"

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
