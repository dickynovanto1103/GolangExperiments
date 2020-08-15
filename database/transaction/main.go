package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/testDB")
	if err != nil {
		log.Fatal("error in connecting to DB:", err)
	}
	//db.SetConnMaxLifetime()
	db.SetMaxOpenConns(100)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("error in db PING, err: %v", err)
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFn()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("error in beginning transaction, err: %v", err)
	}
	_, err = tx.Exec("INSERT INTO Person(Name, Age) VALUES(\"dicky\", 12)")
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Fatalf("error rollback, err: %v", errRollback)
		}
		log.Printf("rollback success for INSERT INTO Person")
		return
	}

	_, err = tx.Exec("INSERT INTO Person(Name, Age) VALUES(\"dicky\", 13)")
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Fatalf("error rollback, err: %v", errRollback)
		}
		log.Printf("rollback success for INSERT INTO Man")
		return
	}

	err = tx.Commit()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Fatalf("error rollback, err: %v", errRollback)
		}
		log.Printf("rollback success for Commit")
		return
	}
}
