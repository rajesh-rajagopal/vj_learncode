package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)


var (
	httpPort = "2049"
	listenIP = "0.0.0.0"
)

func main() {

	server := NewHttpServer(listenIP, httpPort)
	server.Serve()
}


type Server struct {
	ListenHost string
	ListenPort string
	Logger     *log.Logger
}

func NewHttpServer(listenIP, httpPort string) *Server {

	res := Server{
		ListenHost: listenIP,
		ListenPort: httpPort,
		Logger:     log.New(os.Stdout, "server> ", log.Ltime|log.Lshortfile)}

	http.HandleFunc("/", res.HandleIndex)

	return &res
}

func (s *Server) Serve() {

	listenString := s.ListenHost + ":" + s.ListenPort
	s.Logger.Println("Serving http://" + listenString)
	s.Logger.Fatal(http.ListenAndServe(listenString, nil))
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {

	defer timeTrack(time.Now(), w, "HandleIndex")

	s.httpHeaders(w)
        w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")
}

func (s *Server) httpHeaders(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
}

func timeTrack(start time.Time, w http.ResponseWriter, name string) {

	elapsed := time.Since(start)
	io.WriteString(w, "<footer>"+name+" generated in "+elapsed.String()+"</footer>")
}
