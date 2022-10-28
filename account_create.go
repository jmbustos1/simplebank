package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	db "github.com/jmbustos1/simplebank/db/sqlc"
)

const (
	dbDriver2 = "postgres"
	dbSource2 = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

// dbSource = "postgresql://root:secret@localhost:54
)

func main3() {
	//var q *db.Store
	//var d *sql.DB

	var err error
	conn, err := sql.Open(dbDriver2, dbSource2)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	arg := db.CreateAccountParams{
		Owner:    "mmmebb",
		Currency: "USD",
		Balance:  0,
	}

	account, err := store.CreateAccount(context.Background(), arg)
	if err != nil {
		fmt.Println("Hello World!")
		return
	} else {
		fmt.Println("New Account added", account)
	}
	//q = db.New(d)
}
