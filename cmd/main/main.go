package main

import (
	"log"
	"net/http"

	"github.com/balajiss36/go-bookstore/pkg/routes" // Will ttell that the routes reside in the bookstore-route file
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // This variable is created in the routes folder, check the code and the request r is the router as mentioned in the function in that folder.
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // Based on the request r as per the routes, the function of controller is handled.
}

// So here http listens as 9010 port, the request from user is sent to r from there to the routes folder.
// There, based on the type of the request(GET, PUT, DEL) the controller function is handled.
// Flow of code then goes from routes to controller and then to the models where the function is executed and interacting with the db.
