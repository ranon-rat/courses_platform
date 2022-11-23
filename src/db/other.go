package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"crypto/sha256"
	"encoding/hex"
)

func Request() (database *sql.DB) {
	var err error

	if database, err = sql.Open("sqlite3", "data/database.db"); err != nil {
		log.Panicln(err)

	}

	return
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "data/database.db")
	if err != nil {
		panic(err)
	}
	return db

}

// sirve para generar un hash unico y no exponerla contraseña y hacer mas complicado el poder hacer fuerza bruta para explotarla

func hashIt(pass string) string {
	return hex.EncodeToString(sha256.New().Sum([]byte(pass)))
}

func Hash(s string) string {
	return hashIt(s)
}
