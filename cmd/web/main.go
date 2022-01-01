package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/weigoldj/bookings/pkg/config"
	"github.com/weigoldj/bookings/pkg/handlers"
	"github.com/weigoldj/bookings/pkg/render"
)

// constants that i dont want to change
const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

//const template = "templates/"

func main() {

	// change this to true
	app.InProduction = false

	// setup session
	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("FATAL: Cannot create template cache ", err)
	}

	// set config values
	app.TemplateCache = tc
	app.UseCache = true

	// creaate handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// mux := http.NewServeMux()
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	// Bind to a port and pass our router in
	fmt.Println("starting server, listening on " + portNumber)
	log.Fatal(srv.ListenAndServe())

}
