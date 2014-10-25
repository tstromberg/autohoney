// sqlite_store implements a SQLite based data store.
package sqlite_store

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tstromberg/autohoney/objects"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// SQL statements to create a new database.
var schema = `
CREATE TABLE instances (
	id INTEGER PRIMARY KEY,
	name TEXT,

	provider TEXT,
	provider_id TEXT,

	image TEXT,
	recipes TEXT,

	creation_time TIMESTAMP,
	start_time TIMESTAMP,
	stop_time TIMESTAMP
)
`

// NewStore creates and returns a storage object.
func NewStore(path string) (*Store, error) {
	s := Store{Path: path}
	return s.Open()
}

// Store is our storage object.
type Store struct {
	// Path is a filesystem path where the SQLite database will be created.
	Path string
	// db is a sqlx.DB handle.
	db *sqlx.DB
}

// createDB creates a new database at a given path.
func (s *Store) create() error {
	dir := filepath.Dir(s.Path)
	log.Printf("Creating %s ...", dir)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	log.Printf("Opening %s ...", s.Path)
	db, err := sqlx.Open("sqlite3", s.Path)
	if err != nil {
		return err
	}
	log.Print("Executing schema ...")
	_, err = db.Exec(schema)
	if err != nil {
		os.Remove(s.Path)
		return err
	}
	s.db = db
	return nil
}

// Open creates a valid store.
func (s *Store) Open() (*Store, error) {
	log.Printf("Open() called, DBPath is %s", s.Path)

	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		err := s.create()
		if err != nil {
			log.Printf("Problem creating database: %s", err)
			return nil, err
		}
		return s, err
	}
	log.Printf("Opening %s ...", s.Path)
	db, err := sqlx.Open("sqlite3", s.Path)
	if err != nil {
		log.Printf("Problem opening database: %s", err)
		return nil, err
	}
	s.db = db
	return s, nil
}

// queryInstance queries the instance table for an id (or 0 for everything).
func (s *Store) QueryInstances(q objects.InstanceQuery) (instances []objects.Instance, err error) {
	instance := objects.FlatInstance{}
	query := "SELECT id, name, image, recipes FROM instances"
	var rows *sqlx.Rows

	if q.Id != 0 {
		rows, err = s.db.Queryx(query+" WHERE id=:id", q.Id)
	} else {
		rows, err = s.db.Queryx(query)
	}
	if err != nil {
		return
	}
	for rows.Next() {
		rows.StructScan(&instance)
		instances = append(instances, instance.Instance())
	}
	return
}

// AddInstance adds an instance.
func (s *Store) AddInstance(i objects.Instance) error {
	log.Printf("AddInstance %v", i)
	result, err := s.db.NamedExec(`INSERT INTO instances (name, image) VALUES (:name, :image)`, i)
	log.Printf("result: %V", result)
	return err
}

// DeleteInstance deletes an instance.
func (s *Store) DeleteInstance(i objects.Instance) error {
	log.Printf("DeleteInstance %v", i)
	_, err := s.db.NamedExec(`DELETE FROM instances WHERE id=:id`, i)
	return err
}
