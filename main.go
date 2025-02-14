package main

import (
	"database/sql"
	"horizon/deploy"
	"horizon/server"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=horizon-dev password=horizon-dev sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to the database - ", err)
	}

	env := server.NewEnv(db)
	app := echo.New()

	s := server.NewServer(env, app)

	s.RegisterRoutes(
		deploy.RegisterRoutes,
	)

	err = s.Start(":8080")
	if err != nil {
		log.Fatal("failed to start the server - ", err)
	}
}
