package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xyztavo/go-gym/internal/configs"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/routes"
)

func main() {
	HandleArgs()
	r := routes.SetupRoutes()
	fmt.Printf("\n\x1b[32m%s\x1b[0m\n\x1b[33mhttp://localhost%s\x1b[0m", "Ｇｏ Ｇｙｍ 💪", configs.GetPort())
	http.ListenAndServe(configs.GetPort(), r)
}

func HandleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "migrate":
			if err := database.Migrate(); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("🚢 | Database Migrated!")
				os.Exit(0)
			}
		case "seed":
			if err := database.Seed(); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("🌱 | Database Seeded!")
				os.Exit(0)
			}
		}
	}
}
