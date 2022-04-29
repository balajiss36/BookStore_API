package models

import (
	"github.com/balajiss36/go-bookstore/pkg/config" // To connect with the database.
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model         // Structure to store in the database
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()        // To access the Connect function in the app.go file created in config folder
	db = config.GetDB()     // Access the db value from the GetDB function in config folder, we need to store it locally
	db.AutoMigrate(&Book{}) // AutoMigrate will create tables as per the db schema and for the Book struct
	// db helps to talk to the database

}

// Route send request to the Controller, controller will access the model file and here we need to create function for each of the API request as per the route file.
// Model code talks to the database to make changes and sends data output back to main.go

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // Newrecord function exists in the GORM which writes mysql queries to get data/ change data from db
	db.Create(&b)
	return b // We first create a new record for b and then write it to record using Create function and return the value back to the Book struct
}

func GetAllBooks() []Book {
	var Books []Book // Create a var for struct Books because we cannot directly pass a struct to the db
	db.Find(&Books)  // Since it is a struct, we need to use pointers to find the location of the Struct in the db
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) { // Point to the mem address to get the value
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook) // Where is in the db function of gorm, based on the ID passed and find it.
	return &getBook, db                       // Just return the book as per the id from the db & and db variable
}

func DeleteBook(ID int64) Book { // To delete the value, no pointer just pass by value
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book

}
