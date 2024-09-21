package main

import (
	"log"
	"net/http"

	"github.com/xyztavo/go-gym/internal/configs"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/routes"
)

func main() {
	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}
	r := routes.SetupRoutes()
	http.ListenAndServe(configs.GetPort(), r)
}
