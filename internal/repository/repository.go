package repository

import (
	"go-crud-api/internal/models"
)

type DataBaseRepo interface {
	Connect() error
	CheckDatabase(initScript string) error
	Close()

	GetDatabyID(id int) (*models.Data, error)
	InsertData(data models.Data) int
	DeleteDatabyID(id int) (int, error)
	UpdateData(data models.Data) (int, error)
}
