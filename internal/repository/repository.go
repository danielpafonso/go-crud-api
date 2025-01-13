package repository

import (
	"go-crud-api/internal/models"
)

type DataBaseRepo interface {
	Connect() error
	CheckDatabase() error

	GetDatabyID(id int) (*models.Data, error)
	InsertData(data models.Data) (int, error)
	DeleteDatabyID(id int) error
	UpdateData(data models.Data) error
}
