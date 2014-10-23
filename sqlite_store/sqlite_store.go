// sqlite_store implements a SQLite based data store.
package sqlite_store

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tstromberg/autohoney/objects"

	"github.com/jmoiron/sqlx"
)

// SQL statements to create a new database.
var schema = `
CREATE TABLE instances (
	id int,
	uuid text,
	name text,
	image text,
	recipes text,
	creation_time timestamp,
	start_time timestamp,
	first_timestampercept_time timestamp,
	end_time timestamp
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

func (s *Store) ListInstances() (instances []objects.Instance, err error) {
	log.Printf("ListInstances ...")
	instance := objects.SavedInstance{}
	rows, err := s.db.Queryx("SELECT * FROM instances")
	if err != nil {
		return
	}
	for rows.Next() {
		rows.StructScan(&instance)
		instances = append(instances, instance.Instance())
	}
	log.Printf("ListInstances return: %v", instances)
	return
}

func (b *Store) AddInstance(i objects.Instance) error {
	log.Printf("AddInstances %v", i)
	return nil
}

func (b *Store) UpdateInstance(i objects.Instance) error {
	log.Printf("UpdateInstances %v", i)
	return nil
}

func (b *Store) DeleteInstance(i objects.Instance) error {
	log.Printf("DeleteInstance %v", i)
	return nil
}
