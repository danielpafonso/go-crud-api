package main

import (
	"encoding/json"
	"log"
	"net/http"

	// "flag"

	"go-crud-api/api"
	"go-crud-api/api/middleware"
)

func handlerGeneral(w http.ResponseWriter, r *http.Request) {
	log.Println("General response")
	w.Write([]byte("Hello, you hit our general endpoint"))
}

func main() {
	// Add flags
	//    port
	router := http.NewServeMux()

	apiHandler := api.Handler{
		Data: api.LoadMapData(),
	}
	apiHandler.NextID = len(apiHandler.Data) + 1

	// generic handler
	router.HandleFunc("/", handlerGeneral)
	router.HandleFunc("GET /data/{id}", apiHandler.FindByID)
	router.HandleFunc("POST /data", apiHandler.InsertData)
	router.HandleFunc("DELETE /data/{id}", apiHandler.DeleteByID)
	router.HandleFunc("PUT /data/{id}", apiHandler.UpdateByID)

	router.HandleFunc("/coffee", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(http.StatusText(http.StatusTeapot)))
	})

	// debug
	router.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(apiHandler.Data)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}
	log.Println("Server listing in port :8080")
	server.ListenAndServe()
}
