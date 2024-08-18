package main

import (
	"fmt"
	"log"

	"github.com/isaiaspereira307/gowallet/configs"
	"github.com/isaiaspereira307/gowallet/internal/db"
	"github.com/isaiaspereira307/gowallet/routes"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatal("Failed to load configurations:", err)
		panic(err)
	}
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close()

	queries := db.New(conn)
	port := fmt.Sprintf(":%s", configs.GetServerPort())
	r := routes.SetupRouter(queries)
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
