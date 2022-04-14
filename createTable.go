package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func TableCreation() {
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user 
	(
		id INTEGER PRIMARY KEY, 
		uid TEXT, 
		username TEXT, 
		email TEXT, 
		passwd TEXT,
		avatar TEXT

	);`)

	Debug(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts 
	(
		pid INTEGER PRIMARY KEY, 
		content TEXT,
		like INTEGER, 
		category TEXT,
		uid TEXT
	);`)

	Debug(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments 
	(
		cid INTEGER PRIMARY KEY,
		pid INTEGER,
		like INTEGER, 
		content TEXT,
		uid TEXT
	); `)

	Debug(err)
}
