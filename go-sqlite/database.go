package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // database pattern
)

func openDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, fmt.Errorf("Path %s: %s", path, err)
	}

	return db, nil
}

func createDemoTable(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Begin: %s", err)
	}

	statement, err := tx.Prepare(`
    CREATE TABLE demo (
      id INTEGER PRIMARY KEY,
      col1 VARCHAR(255)
    )
  `)
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Commit: %s", err)
	}

	return nil
}
