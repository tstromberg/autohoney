// autohoneyd is the autohoney daemon.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tstromberg/autohoney/sqlite_store"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// Path where data should be stored to. Currently only SQLite is supported.
	StoragePath = os.ExpandEnv("$HOME/.autohoney/sqlite.db")

	// Global storage object
	Store *sqlite_store.Store

	// Address to listen on.
	ListenAddress = ":8181"
)

// hello world, the web server
func InstancesHandler(w http.ResponseWriter, req *http.Request) {
	var response []byte

	instances, err := Store.ListInstances()
	if err == nil {
		response, err = json.Marshal(instances)
	}
	if err != nil {
		log.Print(err)
		response = []byte(fmt.Sprintf(`{ "error": "%s" }`, err))
	}
	w.Write(response)
}

func main() {
	// Get the storage object setup.
	s, err := sqlite_store.NewStore(StoragePath)
	if err != nil {
		log.Fatal(err)
	}
	Store = s

	http.HandleFunc("/instances", InstancesHandler)

	log.Printf("Listening at %s ...", ListenAddress)

	err = http.ListenAndServe(ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
