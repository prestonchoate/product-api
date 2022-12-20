package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Base is a simple handler
type Base struct {
	l *log.Logger
}

// NewBase creates a new Base handler with the given logger
func NewBase(l *log.Logger) *Base {
	return &Base{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Base) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle requests")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Base %s", b)
}