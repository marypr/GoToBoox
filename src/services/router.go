package services

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/authentification/midlware"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

const (
	apiRoute = "/api/v1"
)

type UserService struct {
	UsersRepo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		UsersRepo: repository,
	}
}

var router *gin.Engine

//Start is a function that starts server and initializes all the routes.
func Start(port string) {
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		isLoggedIn := midlware.CheckLoggedIn(c)
		if !isLoggedIn {
			guest := true
			c.HTML(http.StatusOK, "index.tmpl.html", guest)
		} else {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
		}
	})

	router.GET(apiRoute, IndexHandler)
	initUserProfileRoutes()
	initBooksRoutes()
	initCategoriesRoutes()
	router.Run(":" + port)
}

func initUserProfileRoutes() {

	// Use the SetUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(midlware.SetUserStatus())

	service := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	userRoutes := router.Group(apiRoute)
	{
		// Handle POST requests at /api/v1/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", midlware.EnsureNotLoggedIn(), service.PerformLoginHandler)

		// Handle GET requests at /api/v1/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", midlware.EnsureLoggedIn(), service.LogoutHandler)

		// Handle POST requests at /api/v1/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", midlware.EnsureNotLoggedIn(), service.UserCreateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/userProfile", midlware.EnsureLoggedIn(), service.UserGetHandler)

		// Handle the GET requests at /api/v1/register
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.PUT("/userProfile", midlware.EnsureLoggedIn(), service.UserUpdateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.DELETE("/userProfile", midlware.EnsureLoggedIn(), service.UserDeleteHandler)
	}

	// Show the login page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/login", midlware.EnsureNotLoggedIn(), ShowLoginPage)

	// Show the registration page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/register", midlware.EnsureNotLoggedIn(), ShowRegistrPage)

	// Show the user's profile page
	// Ensure that the user is logged in by using the middleware
	router.GET("/usersProfile", midlware.EnsureLoggedIn(), ShowUsersProfilePage)
}

func initBooksRoutes() {
	bookService := BookService{postgres.NewBooksRepository(dataBase.Connection)}
	//get all books in certain category
	router.GET("categories/:cat_id/books", bookService.getBooks)
	//get all books
	router.GET("/books", bookService.showAllBooks)
	//get books by it's ID
	router.GET("categories/:cat_id/book/:book_id", bookService.getBook)
	router.GET("/mostPopularBooks", bookService.FiveMostPop)

}

func initCategoriesRoutes() {
	categoriesService := NewCategoriesService(postgres.CategoryRepoPq{})
	{
		router.GET("/categories", categoriesService.AllCategories)
	}
}

