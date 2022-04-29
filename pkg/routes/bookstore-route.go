package routes

import (
	"github.com/balajiss36/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST") // If User hits the /book path in the API, it should go to the
	// controller folder in pkg(that's why we have imported the pkg repo from the same project) and launch the createBook function

	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET") // If the method is GET for /book path, launch the GetBook function
	router.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
