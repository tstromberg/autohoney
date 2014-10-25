// autohoneyd is the autohoney daemon.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	// This over-engineered store is designed to be replaceable.
	"github.com/tstromberg/autohoney/objects"
	store "github.com/tstromberg/autohoney/sqlite_store"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var (
	// Path where data should be stored to. Currently only SQLite is supported.
	StoragePath = os.ExpandEnv("$HOME/.autohoney/sqlite.db")

	// Global storage object
	Store *store.Store

	// Address to listen on.
	ListenAddress = ":8181"
)

// ListInstances handles GET /instances/ and /instances/:id
func ListInstances(c web.C, w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	q := objects.InstanceQuery{}
	if err != nil {
		q.Id = id
	}

	var response []byte
	instances, err := Store.QueryInstances(q)
	if err == nil {
		response, err = json.Marshal(instances)
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// AddInstance handles POST /instances/
func AddInstance(c web.C, w http.ResponseWriter, req *http.Request) {
	var i objects.Instance
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if i.Name == "" {
		http.Error(w, "Missing required field: name", http.StatusBadRequest)
		return
	}
	if i.Image == "" {
		http.Error(w, "Missing required field: image", http.StatusBadRequest)
		return
	}

	err = Store.AddInstance(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, "Everything is awesome", http.StatusCreated)
}

func main() {
	// Get the storage object setup.
	s, err := store.NewStore(StoragePath)
	if err != nil {
		log.Fatal(err)
	}
	Store = s

	log.Printf("Listening at %s ...", ListenAddress)
	goji.Get("/instances", ListInstances)
	goji.Get("/instances/:id", ListInstances)
	goji.Post("/instances", AddInstance)

	goji.Serve()
}
