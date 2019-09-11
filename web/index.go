// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// Index is the front page of the web application
func (srv Web) Index(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(srv.Settings.DocRoot, srv.Settings.IndexTemplate)

	t, err := template.ParseFiles(p)
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
