package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kseniadobrovolskaia/task-9/internal/config"
)

type DataBase struct {
	Db *sql.DB
}

func NewDataBase() *DataBase {
	return &DataBase{}
}

func (bd *DataBase) Open(cfg *config.Config) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPswrd, cfg.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return errors.New("error while open: " + err.Error())
	}
	bd.Db = db

	err = bd.Db.Ping()
	if err != nil {
		return errors.New("error while ping: " + err.Error())
	}
	return nil
}

func (bd *DataBase) Close() {
	bd.Db.Close()
}
