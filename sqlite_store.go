// sqlite_store implements a SQLite based data store.
package sqlite_store

import (
	"os"
	"path/filepath"

	"objects"

	"github.com/jmoiron/sqlx"
)

// SQL statements to create a new database.
var Schema = `
CREATE TABLE instances (
	id int,
	uuid text,
	name text,
	image text,
	recipes text,
	creation_time timestamp,
	start_time timestamp,
	first_timestampercept_time timestamp,
	end_time timestamp,
)
`

// NewStore creates and returns a storage object.
func NewStore(string path) (Store, error) {
	s = Store{Path: path}
	return s.Open()
}

// Store is our storage object.
type Store struct {
	// Path is a filesystem path where the SQLite database will be created.
	Path string
	// db is a sqlx.DB handle.
	db nil
}

// createDB creates a new database at a given path.
func (s *Store) create() error {
	if err := os.MkdirAll(filepath.Dir(s.Path)); err != nil {
		return err
	}

	s.db, err = sqlx.Open("sqlite3", b.Path)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(sql_store.schema)
	if err != nil {
		return err
	}
}

// Open creates a valid store.
func (s *Store) Open() (*Store, error) {
	if _, err := os.Stat(b.Path); os.IsNotExist(err) {
		err := s.create()
		return s, err
	}

	s.db, err = sqlx.Open("sqlite3", b.Path)
	return s, err
}

func (b *Store) ListInstances() (instances []objects.Instance, err error) {
	instance := objects.SavedInstance{}
	rows, err := db.Queryx("SELECT * FROM instances")
	for rows.Next() {
		rows.Structscan(&instance)
		instances = append(instances, instance.Instance())
	}
}

func (b *Store) AddInstance(i objects.Instance) error {
}

func (b *Store) UpdateInstance(i objects.Instance) error {
}

func (b *Store) DeleteInstance(i objects.Instance) error {
}
