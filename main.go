package main

import (
	"errors"
	"log"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func main() {
	conn, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	mock.ExpectBegin()
	want := errors.New("STMT ERROR")
	mock.ExpectPrepare("SELECT").WillReturnCloseError(want)

	txn, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare("SELECT")
	if err != nil {
		log.Fatal(err)
	}

	if err := stmt.Close(); err != want {
		log.Fatalf("Got = %v, want = %v", err, want)
	}
}
