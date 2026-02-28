pakage main


////we will be working on an api that will be able to 
////store ,retrieve and change data 


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
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

func main() {
	router := gin.Default()
	router.Get("/books",getbooks)
	get.Run("localhost:8989")
}


