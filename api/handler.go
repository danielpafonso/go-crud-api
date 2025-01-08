package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

// handlers code
type Handler struct {
	Data   map[int]Data
	NextID int
}

func (hdl *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	requestID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	data, ok := hdl.Data[requestID]

	if !ok {
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

	hdl.Data[hdl.NextID] = Data{
		ID:    hdl.NextID,
		Value: string(body),
	}
	log.Println("Written data", hdl.Data[hdl.NextID])
	hdl.NextID += 1
}

func (hdl *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	requestID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	if _, ok := hdl.Data[requestID]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}
	// delete map entry
	delete(hdl.Data, requestID)
	log.Println("Delete entry with id:", requestID)
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

	// check if data exists
	data, ok := hdl.Data[requestID]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Data not found"))
		return
	}

	// get Body data
	body, _ := io.ReadAll(r.Body)
	// change data
	data.Value = string(body)
	hdl.Data[requestID] = data
	log.Println("Updated data:", data)
}
