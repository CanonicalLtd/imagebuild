package web

import (
	"fmt"
	"launchpad.net/ce-web/imagebuild/config"
	"net/http"
)

type Web struct {
	Settings *config.Settings
}

// NewWebService starts a new web service
func NewWebService(settings *config.Settings) *Web {
	return &Web{
		Settings: settings,
	}
}

func (srv Web) Start() error {
	fmt.Printf("Starting service on port %s\n", srv.Settings.Port)
	srv.Router()
	return http.ListenAndServe(srv.Settings.Port, nil)
}

// Router returns the application router
func (srv Web) Router() {
	// Start the web service router
	//router := mux.NewRouter()

	//router.Handle("/api/login", Middleware(http.HandlerFunc(srv.Login))).Methods("POST")

	// Serve the static path
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Settings.DocRoot+"/static")))
	//router.PathPrefix("/static/").Handler(fs)
	http.Handle("/static/", fs)

	// Default path is the index page
	//router.Handle("/", Middleware(http.HandlerFunc(srv.Index))).Methods("GET")
	//router.NotFoundHandler = Middleware(http.HandlerFunc(srv.Index))
	http.Handle("/", Middleware(http.HandlerFunc(srv.Index)))
}
