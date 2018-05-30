package cmf

import (
	"context"
	"database/sql"
	"net/http"
)

// Amaranth manages the application state.
type Amaranth struct {
	dir string // The instance folder.
	*config
	db     sql.DB
	server http.Server
}

// New initializes a new amaranth instance using the configuration details in
// the passed directory.
func New(dir string) (*Amaranth, error) {
	// The instance directory needs to be stored as an absolute path, in case
	// the process changes its working directory later.
	dir = filepath.Abs(dir)
	// Check for and create lock file.

	conf, err := readConfig(dir)
	if err != nil {
		return nil, err
	}
	db, err := connectDB(conf.Database)
	if err != nil {
		return nil, err
	}
	return &Amaranth{dir, conf, db}, nil
}

func (a *Amaranth) Reload() error {
	conf, err := ReadConfig(a.dir)
	if err != nil {
		return err
	}
	a.conf = conf
	return nil
}

func (a *Amaranth) Stop() error {
	return nil
}
