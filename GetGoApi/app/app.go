package app

import (
	"GetGoApi/database"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func New() *App {

	a := &App{

		Router: mux.NewRouter(),
	}
	a.initRoutes()
	return a

}

func (a *App) initRoutes() {

	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/app/orders", a.GetOrderHandler()).Methods("GET")

}
