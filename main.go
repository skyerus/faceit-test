package main

import (
	"log"
	"os"

	"github.com/skyerus/faceit-test/pkg/api"
	"github.com/skyerus/faceit-test/pkg/db"
	"github.com/skyerus/faceit-test/pkg/env"
)

func main() {
	env.SetEnv()
	main := &api.App{}
	conn, err := db.OpenDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	if os.Getenv("INIT_TABLES") == "true" {
		err := db.InitiateTables(conn)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	main.Initialize(conn)
	main.Run(":80")
}
