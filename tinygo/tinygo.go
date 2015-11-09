package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var (
	defaultString = "Hello there\n"
)

func main() {

	s := &http.Server{
		Addr:    "tcp://0.0.0.0:80",
		Handler: mux(http.DefaultServeMux),
	}

	if len(os.Args) > 1 {
		defaultString = os.Args[1] + "\n"
	}

	http.HandleFunc("/", defaultResponse)
	l, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		fatal(err)
	}
	if err := s.Serve(l); err != nil {
		fatal(err)
	}
}

func mux(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}

func defaultResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, defaultString)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
