pakage main


////we will be working on an api that will be able to 
////store ,retrieve and change data 


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)


type book struct {
	ID string
	Title string
	Author string
	Quantity int
}