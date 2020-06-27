package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver

	"github.com/gorilla/mux"
)

// App - Root component
type App struct {
	Router *mux.Router
}

type router struct {
	conn *sql.DB
}

func newRouter(conn *sql.DB) *router {
	return &router{conn}
}

// Initialize - Initialize app
func (a *App) Initialize(conn *sql.DB) {
	router := newRouter(conn)
	a.Router = mux.NewRouter()
	a.setRouters(router)
}

func (a *App) setRouters(router *router) {
	a.Router.HandleFunc("/", healthCheck).Methods("GET", "OPTIONS")
}

// Run - Run the app
func (a *App) Run(host string) {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         host,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  18 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
