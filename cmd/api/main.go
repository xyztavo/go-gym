package main

import (
	"fmt"
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
	fmt.Printf(`

┏┓  ┏┓     
┃┓┏┓┃┓┓┏┏┳┓
┗┛┗┛┗┛┗┫┛┗┗
       ┛    
http://localhost%v
	`, configs.GetPort())
	http.ListenAndServe(configs.GetPort(), r)
}
