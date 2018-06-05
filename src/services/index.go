package services

import (
	"github.com/gin-gonic/gin"
	//"github.com/metalscreame/GoToBoox/src/models"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

func IndexHandler(c *gin.Context) {
	type Categories struct{
		ID int
		Title string
	}

	type PopularBooks []struct{
		Id int64
		Title string
		Description string
		Popularity float64
		CategoryId int64
	}

	type Data struct{
		cats []repository.Categories
		PopularBooks
	}

	categories, err := postgres.CategoryRepo.GetAllCategories()
	if err != nil{

	}
	popularBooks := PopularBooks{
		{0, "Title-1", "Description", 5.0, 0},
		{1, "Title-2", "Description", 4.5, 1},
		{2, "Title-3", "Description", 4.2, 2},
	}


	output := Data{categories, popularBooks}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}