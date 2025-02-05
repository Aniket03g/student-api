package sqlite

import (
	"database/sql"

	"github.com/Aniket03g/students-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

//var _ storage.Storage = (*Sqlite)(nil)

// added the _ above are using the sqlite indirectly
type Sqlite struct {
	Db *sql.DB
}

// creating sqlite instance
func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	//used backticks `` here for multiline formating
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}

// receiver func to implement interface, to make the db as plug and play dependency
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastid, nil

}
