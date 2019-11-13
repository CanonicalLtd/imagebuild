// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

// StoreSearchHandler fetches the available snaps from the store
func (srv Web) StoreSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)

	// Decode the request body
	bld, err := decodeBuildRequest(w, r)
	if err != nil {
		return
	}

	// Set the arch header as a search filter
	meta := srv.BoardSrv.BoardMeta(bld.BoardID, bld.OSID)
	headers := map[string]string{
		"X-Ubuntu-Architecture": meta.Arch,
	}

	// Set up the search URL
	vars := mux.Vars(r)
	url := srv.Settings.StoreURL + "snaps/search?q=" + vars["snapName"]

	// Set search options for core and classic
	if strings.HasPrefix(bld.OSID, "core") {
		url = url + "&confinement=strict"
	}
	if strings.HasPrefix(bld.OSID, "classic") {
		url = url + "&confinement=strict,classic"
	}

	resp, err := get(url, headers)
	if err != nil {
		fmt.Fprint(w, "{}")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(w, "{}")
		return
	}

	fmt.Fprint(w, string(body))
}

var get = func(u string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil
	}

	// Set headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return client.Do(req)
}
