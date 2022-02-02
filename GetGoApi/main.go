package main

import (
	"GetGoApi/app"
	"GetGoApi/database"
	"log"
	"net/http"
	"os"
)

func main() {

	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	err = http.ListenAndServe(":9000", nil)
	check(err)

}

func check(e error) {

	if e != nil {
		log.Println(e)
		os.Exit(1)

	}
}
