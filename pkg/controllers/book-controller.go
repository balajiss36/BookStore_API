package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/balajiss36/go-bookstore/pkg/models"
	"github.com/balajiss36/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// Input point is to convert the output to JSON to sent it to API back
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() // If this function is invoked, it will check the models folder and check the GetAllBooks function and return the value
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // Response sent back to API is first Marshalled into JSON using res, the input for JSON will the o/p we received from the GetALlBooks function in the models folder.
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // To access the book ID inside the request r, it will be passed in the request
	bookID := vars["bookID"]                  // bookID will be there in URL so we don't need to Unmarshall it JSON
	ID, err := strconv.ParseInt(bookID, 0, 0) // Convert the string value of Book ID into Int so that it can read by the function.
	if err != nil {
		fmt.Println("Error while accessing ID")
	}
	bookDetails, _ := models.GetBookByID(ID) // the converted value is sent to Function Getbookby ID in models. Blank is because we donot need to use the db value here as mentioned in the function.
	res, _ := json.Marshal(bookDetails)      // Assign the value we get and convert it to JSON and send it back
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Here User also inputs the data so we need to parse it from JSON to Go readable, invoke the function and then Marshall it back to JSON
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}   // This variable will be pointing to the Book struct in the models folder
	utils.ParseBody(r, CreateBook) // We need to Parse the body of the request "r" from user by Unmarshalling it using the utils ParseBody function so that the db will understand
	// This parsed value will be sent to the CreateBookx variable
	b := CreateBook.CreateBook() // This input will be executed with function CreateBook in the models folder
	// CreateBook function will create a record for the user inpurt value and return it back
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID) // After the bookID is converted to String, we run the DeleteBook function in the models function on that ID(input of ID is mentioned there) and assign the value to books
	w.Header().Set("Content-Type", "pkglication/json")
	fmt.Println("Delete Successfull", book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBooks := &models.Book{}   // Create a variable of type Books struct
	utils.ParseBody(r, UpdateBooks) // Unmarshall and send it to updateBooks var so that the request is now as pr the Struct defined.
	vars := mux.Vars(r)             // After unmarshall we need to get the specific value bookID from the request which we are supposed to Update
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0) // Convert it to INT
	if err != nil {
		fmt.Println("Unable to Parse")

	} // We first find the particular ID as per the input and update the values as per request
	bookDetails, db := models.GetBookByID(ID) // First we will get the details of particular Book searched by ID from the Getbookby ID func
	// The value we get from this function is sent to bookdetails variable
	if UpdateBooks.Name != "" { // If the value is not nil, then update it with the request sent in UpdateBooks which is parsee value of r.
		bookDetails.Name = UpdateBooks.Name
	}
	if UpdateBooks.Author != "" {
		bookDetails.Author = UpdateBooks.Author
	}
	if UpdateBooks.Publication != "" {
		bookDetails.Publication = UpdateBooks.Publication
	} // All the updated values are in the bookdetails var
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
