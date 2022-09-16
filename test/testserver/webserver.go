package testserver

import (
	"fmt"
	"net/http"
	"path"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(root string) *Server {
	fileServer := http.FileServer(http.Dir(root))
	return &Server{
		httpServer: &http.Server{
			Addr:    "127.0.0.1:8000",
			Handler: fileServer,
		},
	}
}

func (w *Server) Start() {
	go w.httpServer.ListenAndServe()
}

func (w *Server) Stop() error {
	return w.httpServer.Close()
}

func (w *Server) Address() string {
	return w.httpServer.Addr
}

func (w *Server) WhereIs(file string) string {
	return path.Clean(fmt.Sprintf("http://%s/%s.html", w.Address(), file))
}
