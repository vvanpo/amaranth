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
	"strconv"
	"time"
)

type Config struct {
	Domain   string
	Port     int
	TLS      bool
	TLS_key  string
	TLS_cert string
	// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
	Database string
}

type Server struct {
	conf Config
	db   *sql.DB
	hs   http.Server
}

func (s *Server) Load(conf_path string) (err error) {
	file, err := ioutil.ReadFile(conf_path)
	if err != nil {
		return
	}
	if err = json.Unmarshal(file, &s.conf); err != nil {
		return
	}
	s.db, err = sql.Open("postgres", s.conf.Database)
	if err != nil {
		return
	}
	s.hs.Addr = s.conf.Domain + ":" + strconv.Itoa(s.conf.Port)
	s.hs.Handler = http.NewServeMux()
	return
}

func (s *Server) Start() error {
	if s.conf.TLS {
		return s.hs.ListenAndServeTLS(s.conf.TLS_cert, s.conf.TLS_key)
	}
	return s.hs.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.hs.Shutdown(ctx)
}

func main() {
	var s Server
	var conf_path string
	flag.StringVar(&conf_path, "c", "", "-c [file]")
	flag.Parse()
	if err := s.Load(conf_path); err != nil {
		log.Fatal(err)
	}
	log.Print(s.Start())
}
