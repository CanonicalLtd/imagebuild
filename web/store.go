package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// StoreSearchHandler fetches the available snaps from the store
func (srv Web) StoreSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)

	vars := mux.Vars(r)

	resp, err := get(srv.Settings.StoreURL + "snaps/search?q=" + vars["snapName"])
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

var get = func(u string) (*http.Response, error) {
	return http.Get(u)
}
