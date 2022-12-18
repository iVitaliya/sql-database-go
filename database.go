package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type IDatabase struct {
	// Creates a new table in the database.
	CreateTable func()
	// Adds a number in the database to the provided key-value.
	Add func(key string, value int)
	// Removes all the keys from the database.
	Clear func() bool
	// Deletes the provided key from the database.
	Delete func(key string) bool
	// Retrieves data from the provided key from the database.
	Get func(key string) any
	// Sets a value for the provided key in the database.
	Set func(key string, value any)
	// Substracts a number in the database to the provided key-value.
	Substract func(key string, value int)
}

var (
	rows         *sql.Rows
	result       sql.Result
	stmt         *sql.Stmt
	err          error
	id           int64
	rowsAffected int64
	res          string
)

func getData(db *sql.DB, table string, key string) any {
	rows, err = db.Query(`SELECT value FROM `+table+` WHERE id = ?`, key)
	if err != nil {
		logError("Retrieving Data", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var data string

		err = rows.Scan(&data)
		if err != nil {
			logError("Scanning For Data", err.Error())
		}

		res = data
	}

	return res
}

func setData(db *sql.DB, table string, key string, value any) {
	stmt, err = db.Prepare(`INSERT INTO ` + table + ` WHERE (id, value) VALUES (?, ?)`)
	if err != nil {
		logError("Preparing Setting Key", err.Error())
	}
	defer stmt.Close()

	id, err = result.LastInsertId()
	if err != nil {
		logError("Getting The Last InsertId", err.Error())
	}

	logSetID("AFK", id)
}

func Database(table string) *IDatabase {
	db, err := sql.Open("sqlite3", "./data.sqlite")
	if err != nil {
		logError("Opening Path", err.Error())
	}
	defer db.Close()

	return &IDatabase{
		CreateTable: func() {
			_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ` + table)
			if err != nil {
				logError("Creating A Table", err.Error())
			}
		},
		Add: func(key string, value int) {
			var (
				data      = getData(db, table, key)
				oldNumb   = valueToInt(key, data)
				newNumber = oldNumb + value
			)

			setData(db, table, key, newNumber)
		},
		Clear: func() bool {
			result, err = db.Exec(`DELETE FROM ` + table)
			if err != nil {
				logError("Deleting Keys Failed", err.Error())
				return false
			}

			rowsAffected, err = result.RowsAffected()
			if err != nil {
				logError("Rows Affected Failed", err.Error())
				return false
			}

			logAffectedRows(table, rowsAffected)
			return true
		},
		Delete: func(key string) bool {
			result, err = db.Exec(`DELETE FROM `+table+` WHERE id = ?`, key)
			if err != nil {
				logError("Deleting Key Failed", err.Error())
				return false
			}

			rowsAffected, err = result.RowsAffected()
			if err != nil {
				logError("Rows Affected Failed", err.Error())
				return false
			}

			logAffectedRows(table, rowsAffected)
			return true
		},
		Get: func(key string) any {
			return getData(db, table, key)
		},
		Set: func(key string, value any) {
			setData(db, table, key, value)
		},
		Substract: func(key string, value int) {
			var (
				data      = getData(db, table, key)
				oldNumb   = valueToInt(key, data)
				newNumber = oldNumb - value
			)

			setData(db, table, key, newNumber)
		},
	}
}
