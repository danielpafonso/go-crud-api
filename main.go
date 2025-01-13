package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"go-crud-api/api"
	"go-crud-api/api/middleware"
	"go-crud-api/internal/repository"
	"go-crud-api/internal/repository/dbrepo"
)

func handlerGeneral(w http.ResponseWriter, r *http.Request) {
	log.Println("General response")
	w.Write([]byte("Hello, you hit our general endpoint"))
}

func main() {

	var repo repository.DataBaseRepo

	dbConn := &dbrepo.Sqlite3DB{
		ConnectionString: "data.sqlite",
	}
	repo = dbConn
	err := repo.Connect()
	if err != nil {
		log.Panicln(err)
	}
	err = repo.CheckDatabase()
	if err != nil {
		log.Panicln(err)
	}
	defer dbConn.DB.Close()

	repo.GetDatabyID(2)

	var serverPort string

	flag.StringVar(&serverPort, "p", "8080", "Port which the server will use")
	flag.Parse()

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
		Addr: ":" + serverPort,
		// middleware channing
		Handler: middleware.Logging(rootRouter),
		// Handler: middleware.Logging(middleware.Auth(router)),
	}
	log.Println("Server listing in port :8080")
	server.ListenAndServe()
}
