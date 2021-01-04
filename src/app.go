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
	connect := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	fmt.Println("Up and running!")
}

// Run --> run app
func (a *App) Run(addr string) {}
