package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID string
	Title string
	Author string
	Quantity int
}


var books = []book{
    {ID: "1", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 5},
    {ID: "2", Title: "1984", Author: "George Orwell", Quantity: 3},
    {ID: "3", Title: "Clean Code", Author: "Robert C. Martin", Quantity: 2},
}



func getBooks(c  *gin.Context) {
	c.IndentedJSON(http.StatusOK,books)

}

func createBooks(c *gin.Context) {
	var newBook book
	if err := append(newBook,book)
}



func main(){
	router := gin.Default()
	router.GET("/create",getBooks)
	router.POST("/create",createBooks)
	router.Run("localhost:8055")
}
































// package main

// ////we will be working on an api that will be able to
// ////store ,retrieve and change data

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	//"errors"
// )


// type book struct {
// 	ID string `json:"id"`
// 	Title string `json:"title"`
// 	Author string	`json:"author"`
// 	Quantity int   `json:"quantity"`
// }

// ////a slice of all the books

// var books = []book{
//     {ID: "1", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 5},
//     {ID: "2", Title: "1984", Author: "George Orwell", Quantity: 3},
//     {ID: "3", Title: "Clean Code", Author: "Robert C. Martin", Quantity: 2},
// }

// //
// func getbooks(c *gin.Context){
// 	c.IndentedJSON(http.StatusOK,books)

// }
// //a post req or handler ,note that all the data 
// // that the route will be receiving will be 
// // stored inside the c

// // the use of the c.BindJSON(&newBook)
// // This is very important. it Reads JSON from request body
// //  Matches JSON fields to struct tags 
// // Stores values in newBook
// func createBook(c *gin.Context){
// 	// lets get the book data and create a new var 
// 	// to store the new data into it and should be 
// 	// of the type book
// 	//  &newBook The & means:Pass the memory address so Gin can fill the data.
// 	var newBook book

// 	// error checking
	
// 	if err := c.BindJSON(&newBook); err != nil {
// 		return
// 	}
// 	books =	append(books,newBook)
// 	c.IndentedJSON(http.StatusCreated,newBook)


	
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/books",getbooks)
// 	router.POST("/books",createBook)
// 	router.Run("localhost:8055")
// }



