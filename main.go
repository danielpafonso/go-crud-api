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

	apiHandler := api.Handler{
		Data: api.LoadMapData(),
	}
	apiHandler.NextID = len(apiHandler.Data) + 1

	// generic handler
	rootRouter := http.NewServeMux()
	rootRouter.HandleFunc("/", handlerGeneral)
	rootRouter.HandleFunc("/coffee", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(http.StatusText(http.StatusTeapot)))
	})
	// debug
	rootRouter.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(apiHandler.Data)
	})

	// api router
	apiRouter := http.NewServeMux()
	apiRouter.HandleFunc("GET /{id}", apiHandler.FindByID)

	// api admin router
	apiAdmin := http.NewServeMux()
	apiAdmin.HandleFunc("POST /", apiHandler.InsertData)
	apiAdmin.HandleFunc("DELETE /{id}", apiHandler.DeleteByID)
	apiAdmin.HandleFunc("PUT /{id}", apiHandler.UpdateByID)

	// add auth middleware to api, sending all non defined routes to auth
	apiRouter.Handle("/", middleware.Auth(apiAdmin))

	// add sub routing for Api to root routing
	rootRouter.Handle("/data/", http.StripPrefix("/data", apiRouter))

	server := http.Server{
		Addr: ":8080",
		// middleware channing
		Handler: middleware.Logging(rootRouter),
		// Handler: middleware.Logging(middleware.Auth(router)),
	}
	log.Println("Server listing in port :8080")
	server.ListenAndServe()
}
