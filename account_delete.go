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
	dbDriver3 = "postgres"
	dbSource3 = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

// dbSource = "postgresql://root:secret@localhost:54
)

func main8() {
	//var q *db.Store
	//var d *sql.DB

	var err error
	conn, err := sql.Open(dbDriver3, dbSource3)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	acc := 5
	store := db.NewStore(conn)
	account, err := store.GetAccount(context.Background(), int64(acc))
	if err != nil {
		fmt.Println("No deletion")
		return
	} else {
		fmt.Println("account to be deleted", account)
	}

	store2 := db.NewStore(conn)
	err = store2.DeleteAccount(context.Background(), int64(acc))
	if err != nil {
		fmt.Println("Hello World!")
		return
	} else {
		fmt.Println("New Account deleted", account)
	}
	//q = db.New(d)
}
