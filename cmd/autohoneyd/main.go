// autohoneyd is the autohoney daemon.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"sqlite_store"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// Path where data should be stored to. Currently only SQLite is supported.
	StoragePath = os.ExpandEnv("HOME/.autohoney/sqlite.db")

	// Global storage object
	Store = nil
)

// hello world, the web server
func InstancesHandler(w http.ResponseWriter, req *http.Request) {
	response := json.Marshal(*backend.ListInstances())
	io.WriteString(w, response)
}

func main() {
	// Get the storage object setup.
	s := sqlite_store.Store{Path: StoragePath}
	Store := s.Open()

	http.HandleFunc("/instances", InstancesHandler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
