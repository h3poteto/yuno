package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	db *sql.DB
}

func Initialize() (*database, error) {
	db, err := sql.Open("sqlite3", "./yuno.db")
	if err != nil {
		return nil, err
	}
	d := &database{
		db,
	}
	if err := d.InitializeTable(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *database) InitializeTable() error {
	// Create employee table.
	_, err := d.db.Exec(`
CREATE TABLE IF NOT EXIST employees(
id int(11) NOT NULL AUTO_INCREMENT,
code varchar(255) UNIQUE NOT NULL,
last_name varchar(255),
first_name varchar(255),
key varchar(255) NOT NULL,
type_name varchar(255),
slack_id varchar(255) UNIQUE NOT NULL,
PRIMARY KEY (id))
AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;`)
	if err != nil {
		return err
	}
	return nil
}
