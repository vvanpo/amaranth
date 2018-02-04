package cmf

import (
	"context"
	"database/sql"
	"net/http"
)

// Server manages the application state.
type Server struct {
	dir  string
	*config
	db   sql.DB
	mux  http.Handler
}

// New initializes a new server instance using the configuration details in the
// passed directory.
func New(dir string) (*Server, error) {
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
	s := &Server{dir, conf, db}
	return s, nil
}

func (s *Server) Reload() error {
	conf, err := ReadConfig(s.dir)
	if err != nil {
		return err
	}
	s.conf = conf
	return nil
}

func (s *Server) Stop() error {

}
