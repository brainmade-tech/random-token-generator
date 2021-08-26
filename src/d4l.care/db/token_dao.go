package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var ( // hardcoded config for brevity, otherwise we can externalize it in a config file
	username   = "root"
	password   = ""
	host       = "127.0.0.1"
	port       = "3306"
	dbName     = "d4l_challenge"
	driverName = "mysql"
)

var ( // I made these variables 'global' so they can be accessible from all the functions below
	db   *sql.DB
	stmt *sql.Stmt
	err  error
	tx   *sql.Tx
)

var message = struct {
	connect, beginTx, load, endTx, shutdown string
}{
	"Connection opening error",
	"Transaction begining error",
	"Load file error",
	"Transaction ending error",
	"Connection closing error",
}

func Connect() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)
	db, err = sql.Open(driverName, url)
	if err != nil {
		log.Println(message.connect, "-", err)
		return
	}
	log.Print("Connected to database successfully...")
}

func BeginTX() {
	tx, err = db.Begin()
	if err != nil {
		log.Println(message.beginTx, "-", err)
		return
	}
	log.Print("Begin transaction...")
}

func LoadFromFile(filePath string) { // imports tokens from file to database
	sqlClause := "load data infile '%s' into table tokens (value)"
	sqlClause = fmt.Sprintf(sqlClause, filePath)
	_, err := tx.Exec(sqlClause)
	if err != nil {
		log.Println(message.load, "-", err)
		return
	}
}

// other functions can be added here to support sql clauses like select, update, delete...

func EndTX() {
	err = tx.Commit()
	if err != nil {
		log.Println(message.endTx, "-", err)
		return
	}
	log.Print("End transaction...")
}

func Shutdown() {
	err = db.Close()
	if err != nil {
		log.Println(message.shutdown, "-", err)
		return
	}
	log.Print("Shutdown...")
}
