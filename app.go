package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App : Struct for managing our router and db
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize --> init app
func (a *App) Initialize(user, password, dbname string) {
	fmt.Println(user, password, dbname)
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run --> run app
func (a *App) Run(addr string) {}
