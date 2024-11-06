package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Setup() {
	var err error
	Db, err = sql.Open("mysql", strings.Split(os.Getenv("DATABASE_URL"), "://")[1])
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("____________________")
	fmt.Println("Connected to database!")
	fmt.Println("____________________")
}
