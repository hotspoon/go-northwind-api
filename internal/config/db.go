package config

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"

	_ "modernc.org/sqlite"
)

func SetupDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		log.Fatal().Msgf("failed to connect database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Msgf("unable to reach database: %v", err)
	}

	fmt.Println("Connected to SQLite database!")
	return db
}
