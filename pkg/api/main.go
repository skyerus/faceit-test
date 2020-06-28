package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/skyerus/faceit-test/pkg/cache"

	"github.com/gorilla/mux"
)

// App - Root component
type App struct {
	Router *mux.Router
}

type router struct {
	conn *sql.DB
	c    *cache.Cache
}

func newRouter(conn *sql.DB, c *cache.Cache) *router {
	return &router{conn, c}
}

// Initialize - Initialize app
func (a *App) Initialize(conn *sql.DB, c *cache.Cache) {
	router := newRouter(conn, c)
	a.Router = mux.NewRouter()
	a.setRouters(router)
}

func (a *App) setRouters(router *router) {
	a.Router.HandleFunc("/", healthCheck).Methods("GET")
	a.Router.HandleFunc("/users", router.createUser).Methods("POST")
	a.Router.HandleFunc("/users/{id}", router.getUser).Methods("GET")
	a.Router.HandleFunc("/users/{id}", router.deleteUser).Methods("DELETE")
	a.Router.HandleFunc("/users", router.getUsers).Methods("GET")
	a.Router.HandleFunc("/users/{id}", router.updateUser).Methods("PUT")
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
