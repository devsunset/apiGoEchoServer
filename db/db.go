package db

import (
	"apiServer/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB
var err error

func DbManager() *sql.DB {
	configuration := config.GetConfig()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(configuration.DB_MAXIDLECONNS)
	db.SetMaxOpenConns(configuration.DB_MAXOPENCONNS)

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
