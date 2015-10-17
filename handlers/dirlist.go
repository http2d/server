package handlers

import (
	"fmt"
	"net/http"
)

// DirList class
type DirList struct {
	localPath string
}

// Init initializes the DirList handler
func (dirlist *DirList) Init() error {
	return nil
}

// Reply builds a DirList response
func (dirlist *DirList) Reply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	r.Header.Write(w)
	fmt.Fprintf(w, "This is the dirlist handler")
}

// NewDirList instances a new dirlist handler
func NewDirList() *DirList {
	n := new(DirList)
	return n
}
