package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"go-crud-api/internal/models"
	"go-crud-api/internal/repository"
)

// handlers code
type Handler struct {
	DB     repository.DataBaseRepo
	NextID int
}

func (hdl *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	requestID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	data, err := hdl.DB.GetDatabyID(requestID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Data not found"))
		return
	}

	log.Println("Get data with ID:", requestID)
	// all ok
	json.NewEncoder(w).Encode(data)
}

func (hdl *Handler) InsertData(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	insertedID := hdl.DB.InsertData(models.Data{
		ID:    hdl.NextID,
		Value: string(body),
	})
	log.Println("Written data", insertedID)
}

func (hdl *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	requestID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	numRemoveds, err := hdl.DB.DeleteDatabyID(requestID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}
	log.Println("Delete entry with id:", requestID)
	log.Println("Affected number of rows:", numRemoveds)
}

func (hdl *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	// Get ID
	requestID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
	// check if body have lenght
	bodyLenght := r.Header.Get("Content-Length")
	if bodyLenght == "" || bodyLenght == "0" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// get Body data
	body, _ := io.ReadAll(r.Body)

	data := models.Data{
		ID:    requestID,
		Value: string(body),
	}

	// execute query
	affectedRows, err := hdl.DB.UpdateData(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	if affectedRows == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Data not found"))
		return
	}
	log.Println("Updated data:", data)
}
