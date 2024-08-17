package main

import (
	"database/sql"
	"log"

	"github.com/isaiaspereira307/gowallet/router"

	"github.com/isaiaspereira307/gowallet/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	// Conectando ao banco de dados
	conn, err := sql.Open("postgres", "user=username password=password dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close()

	// Criando um novo objeto Queries
	queries := db.New(conn)

	// Configurando as rotas e passando o objeto Queries
	r := router.SetupRouter(queries)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
