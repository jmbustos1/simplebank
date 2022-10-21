package main

import (
	"database/sql"
	"log"

	"github.com/jmbustos1/simplebank/api"
	db "github.com/jmbustos1/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	//dbSource = "postgresql://root:secret@localhost:54
	serverAdress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
