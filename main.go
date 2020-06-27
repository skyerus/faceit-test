package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/skyerus/faceit-test/pkg/api"
	"github.com/skyerus/faceit-test/pkg/env"
)

func main() {
	env.SetEnv()
	main := &api.App{}
	db, err := openDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	main.Initialize(db)
	main.Run(":80")
}

func openDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_URL")+"?parseTime=true")
	if err != nil {
		return db, err
	}
	return db, nil
}
