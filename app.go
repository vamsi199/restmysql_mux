package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"
)


type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize (user,password,IP,dbname string) {
	connectionString := fmt.Sprintf("%s:%s@(%s:3306)/%s", user, password,IP, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}
func  (a *App) Run(addr string) {}

//The Initialize method is responsible for create a database connection and wire up the routes, and the Run method will simply start the application.
