package server

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"
)

func InitDatabase(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.PGhost, cfg.PGport, cfg.PGuser, cfg.PGname, cfg.PGuser)

	dbHandler, err := sql.Open(cfg.PGdriver, connStr)
	if err != nil {
		log.Fatalf("Error while initiazile db %v", err)
	}
	return dbHandler
}
