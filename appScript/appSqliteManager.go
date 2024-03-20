package appScript

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB() *SQLiteDB {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		fmt.Println(err)
	}
	return &SQLiteDB{db: db}
}

func (s *SQLiteDB) Execute(query string, args ...interface{}) (err error) {
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	return
}

func (s *SQLiteDB) Query(query string, args ...interface{}) *sql.Rows {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func (s *SQLiteDB) QueryRow(query string) *sql.Row {
	row := s.db.QueryRow(query)
	return row
}

func (s *SQLiteDB) Close() {
	err := s.db.Close()
	if err != nil {
		fmt.Println(err)
	}
}
