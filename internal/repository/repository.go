package repository

import (
	"database/sql"
	"go-crud-api/internal/models"
)

type DataBaseRepo interface {
	Connection() *sql.DB

	GetDatabyID(id int) (*models.Data, error)
	InsertData(data models.Data) (int, error)
	DeleteDatabyID(id int) error
	UpdateData(data models.Data) error
}
