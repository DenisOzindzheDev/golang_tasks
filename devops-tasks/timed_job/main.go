package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type LogFile struct {
	File      *os.File
	LogPrefix string
	Separator string
}

type Server struct {
	Server *mux.Router
	LogF   *LogFile
}

func NewServer(log *LogFile) (*Server, error) {
	router := mux.NewRouter()
	return &Server{
		Server: router,
		LogF:   log,
	}, nil
}
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.RequestURI)

		s.LogF.WriteLog(fmt.Sprintf("Server accepted request %v ", r.Method))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
func greetings(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r.Body)
}
func NewLogFile(prefix string, separator string) *LogFile {
	f, _ := os.Create("./log.log")
	return &LogFile{
		File:      f,
		LogPrefix: prefix,
		Separator: separator,
	}
}

func (l *LogFile) WriteLog(log string) {
	l.File.WriteString(l.LogPrefix + log + l.Separator)
}

func main() {
	//init file for log

	LogFile := NewLogFile("[LOGGER]", "\n")
	LogFile.WriteLog("hello world")
	srv, err := NewServer(LogFile)
	if err != nil {
		panic(err)
	}
	srv.Server.HandleFunc("/", greetings)
	srv.Server.Use(srv.loggingMiddleware)
	go func() {
		client := &http.Client{}
		for {
			time.Sleep(time.Second * 1)
			client.Get("http://localhost:8080")
		}
	}()
	go func() {
		client := &http.Client{}
		for {
			time.Sleep(time.Second * 1)
			client.Post("http://localhost:8080", "application/json", nil)
		}
	}()
	err = http.ListenAndServe(":8080", srv.Server)
	if err != nil {
		panic(err)
	}
}
