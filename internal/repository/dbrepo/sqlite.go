package dbrepo

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	"go-crud-api/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

const timeout time.Duration = 3 * time.Second

type Sqlite3DB struct {
	DB               *sql.DB
	ConnectionString string
}

func (db *Sqlite3DB) Connect() error {
	path, _ := os.Executable()
	folder := filepath.Dir(path)
	dbConn, err := sql.Open("sqlite3", filepath.Join(folder, db.ConnectionString))
	if err != nil {
		return err
	}
	db.DB = dbConn
	return nil
}

func (db *Sqlite3DB) Close() {
	db.DB.Close()
}

func (db *Sqlite3DB) CheckDatabase(initScript string) error {
	query := "SELECT name from sqlite_master WHERE type='table' and name='data';"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	row := db.DB.QueryRowContext(ctx, query)

	var table string
	row.Scan(&table)
	if table == "" {
		// create table
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		db.DB.ExecContext(ctx, initScript)
	}
	return nil
}

func (db *Sqlite3DB) GetDatabyID(id int) (*models.Data, error) {
	query := "SELECT * FROM data WHERE id=?;"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	row := db.DB.QueryRowContext(ctx, query, id)

	var data models.Data
	err := row.Scan(&data.ID, &data.Value)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (db *Sqlite3DB) InsertData(data models.Data) int {
	query := "INSERT INTO data (value) values(?);"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := db.DB.ExecContext(ctx, query, data.Value)
	if err != nil {
		log.Panic(err)
	}

	lastID, _ := result.LastInsertId()

	return int(lastID)
}

func (db *Sqlite3DB) DeleteDatabyID(id int) (int, error) {
	query := "DELETE FROM data WHERE id=?;"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := db.DB.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	affected, _ := result.RowsAffected()
	return int(affected), nil
}

func (db *Sqlite3DB) UpdateData(data models.Data) (int, error) {
	query := "UPDATE data SET value=? WHERE id=?;"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := db.DB.ExecContext(ctx, query, data.Value, data.ID)
	if err != nil {
		return 0, err
	}
	affected, _ := result.RowsAffected()
	return int(affected), nil
}
