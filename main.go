package main

import (
	"flag"
	"log"
	"net/http"

	_ "embed"

	"go-crud-api/api"
	"go-crud-api/api/middleware"
	"go-crud-api/internal/repository"
	"go-crud-api/internal/repository/dbrepo"
)

//go:embed init-script.sql
var InitScript string

func handlerGeneral(w http.ResponseWriter, r *http.Request) {
	log.Println("General response")
	w.Write([]byte("Hello, you hit our general endpoint"))
}

func main() {
	var serverPort string
	var dbPath string
	var repo repository.DataBaseRepo

	flag.StringVar(&serverPort, "p", "8080", "Port which the server will use")
	flag.StringVar(&dbPath, "s", "data.sqlite", "Path to the Database file")
	flag.Parse()

	// prep database
	repo = &dbrepo.Sqlite3DB{
		ConnectionString: dbPath,
	}
	err := repo.Connect()
	if err != nil {
		log.Panicln(err)
	}
	err = repo.CheckDatabase(InitScript)
	if err != nil {
		log.Panicln(err)
	}
	defer repo.Close()

	apiHandler := api.Handler{
		DB: repo,
	}

	// generic handler
	rootRouter := http.NewServeMux()
	rootRouter.HandleFunc("/", handlerGeneral)
	rootRouter.HandleFunc("/coffee", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(http.StatusText(http.StatusTeapot)))
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
	log.Println("Server listing in port :" + serverPort)
	server.ListenAndServe()
}
