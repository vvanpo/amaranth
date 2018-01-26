package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type Config struct {
	Domain string
	Port   int
	TLS    bool
	// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
	Database string
}

type Server struct {
	path string
	conf Config
	db   *sql.DB
	hs   http.Server
}

// readConfig parses the config.json file in the instance directory
func readConfig(dir string) (c *Config, err error) {
	file, err := ioutil.ReadFile(filepath.Join(dir, "config.json"))
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &c)
	return
}

func (s *Server) loadConfig(c *Config) (err error) {
	s.conf = *c
	s.db, err = sql.Open("postgres", s.conf.Database)
	if err != nil {
		return
	}
	s.hs.Addr = s.conf.Domain + ":" + strconv.Itoa(s.conf.Port)
	s.hs.Handler = http.NewServeMux()
	return
}

func (s *Server) Start(c *Config) error {
	if err := s.loadConfig(c); err != nil {
		return err
	}
	if s.conf.TLS {
		certFile := filepath.Join(s.path, "certificate.pem")
		keyFile := filepath.Join(s.path, "key.pem")
		return s.hs.ListenAndServeTLS(certFile, keyFile)
	}
	return s.hs.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.hs.Shutdown(ctx)
	s.db.Close()
	return err
}

func main() {
	dir := flag.String("d", "", "instance directory path")
	flag.Parse()
	conf, err := readConfig(*dir)
	if err != nil {
		log.Fatal(err)
	}
	var s Server
	log.Print(s.Start(conf))
}
