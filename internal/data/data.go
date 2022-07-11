package data

import (
	"database/sql"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	DB *sql.DB
}

// internal/data/data.go
func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	//err = MakeMigration(db)
	//if err != nil {
	//	log.Panic(err)
	//}

	data = &Data{
		DB: db,
	}
}

func New() *Data {
	once.Do(initDB)

	return data
}
