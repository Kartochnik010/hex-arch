package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName string, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS arith_history
			(
				id INTEGER PRIMARY KEY,
				date DATE,
				answer INTEGER,
				operation TEXT
			)
	`)
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec()
	if err != nil {
		return nil, err
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDBConnection() error {
	return da.db.Close()
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	// TODO: 00:40
	statement, err := da.db.Prepare("INSERT INTO arith_history (date, answer, operation) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(time.Now(), answer, operation)
	if err != nil {
		return err
	}
	return nil

	// db, err := sql.Open("sqlite3", "./sqlite.db")
	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	// statement, err := db.Prepare("create table if not exists books (id integer primary key, name TEXT, author TEXT)")
	// if err != nil {
	// 	log.Println(err)
	// }
	// statement.Exec()

	// statement, _ = db.Prepare("INSERT INTO books (name, author) VALUES (?, ?)")
	// _, err = statement.Exec("","")
	// rows, _ := db.Query("SELECT id, name, author FROM books")
	// var id int
	// var firstname string
	// var lastname string

	// for rows.Next() {
	// 	rows.Scan(&id, &firstname, &lastname)
	// 	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	// }
	// return nil
}
