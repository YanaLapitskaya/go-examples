package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	host     = "final_project_db_1"
	port     = 5432
	user     = "postgres-dev"
	password = "postres-dev-pass"
	dbname   = "dev"
)

func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")
}
