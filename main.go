package main

////we will be working on an api that will be able to
////store ,retrieve and change data

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)


type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string	`json:"author"`
	Quantity int   `json:"quantity"`
}

////a slice of all the books

var books = []book{
    {ID: "1", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 5},
    {ID: "2", Title: "1984", Author: "George Orwell", Quantity: 3},
    {ID: "3", Title: "Clean Code", Author: "Robert C. Martin", Quantity: 2},
}

//
func getbooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)

}
//a post req or handler ,note that all the data 
// that the route will be receiving will be 
// stored inside the c
func createBook(c *gin.Context){
	var newbook book

	
}

func main() {
	router := gin.Default()
	router.GET("/books",getbooks)
	router.Run("localhost:8989")
}



