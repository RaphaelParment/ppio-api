package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	ppioHTTP "github.com/RaphaelParment/ppio-api/pkg/http"
	"github.com/RaphaelParment/ppio-api/pkg/storage"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, dbTidy, err := storage.SetupDB("ppio")
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbTidy()
	l := log.New(os.Stdout, "ppio :", log.LstdFlags)
	srv := ppioHTTP.NewServer(db, l)

	http.ListenAndServe(":9001", srv)
	return nil
}
