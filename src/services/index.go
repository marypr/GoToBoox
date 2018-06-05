package services

import (
	"github.com/gin-gonic/gin"
	//"github.com/metalscreame/GoToBoox/src/models"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

func IndexHandler(c *gin.Context) {
	type Data struct{
		books []repository.Book
		cats []repository.Categories
	}

	bookRepo := postgres.NewBooksRepository(dataBase.Connection)
	books, _ := bookRepo.GetMostPopularBooks(5)
	catRepo := postgres.CategoryRepoPq{}
	cats, _ := catRepo.GetAllCategories()

	output := Data{books, cats}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}
