package main

import (
	"github.com/3uchris/Book-Management-API-MySQL-Golang/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // from routes.go
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
