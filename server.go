package cms

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type Server struct {
	path string
	conf struct {
		Domain string
		Port   int
		TLS    bool
		// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
		Database string
	}
	db   *sql.DB
	hs   http.Server
}

func NewServer(path string) (s *Server, err error) {
	s = &Server{path: path}
	if err := s.readConfig(); err != nil {
		return nil, err
	}
	if s.conf.TLS {
		certFile := filepath.Join(s.path, "certificate.pem")
		keyFile := filepath.Join(s.path, "key.pem")
		return s, s.hs.ListenAndServeTLS(certFile, keyFile)
	}
	return s, s.hs.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.hs.Shutdown(ctx)
	s.db.Close()
	return err
}

// readConfig parses the config.json file in the instance directory
func (s *Server) readConfig() error {
	file, err := ioutil.ReadFile(filepath.Join(s.path, "config.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, s.conf)
	if err != nil {
		return err
	}
	s.db, err = sql.Open("postgres", s.conf.Database)
	if err != nil {
		return err
	}
	s.hs.Addr = s.conf.Domain + ":" + strconv.Itoa(s.conf.Port)
	s.hs.Handler = http.NewServeMux()
	return nil
}
