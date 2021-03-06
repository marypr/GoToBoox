package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/midlwares"

)

//ShowLoginPage is a handler function that renders static login page
func ShowLoginPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": "Login Page",
			"page": "login",
			"isLoggedIn": isLoggedIn,
		},
	)
}

//ShowRegistrPage is a handler function that renders static registration page
func ShowRegistrPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": "Registration Page",
			"page": "registration",
			"isLoggedIn": isLoggedIn,
		},
	)
}

//UserProfileHandler is a handler func that handle /userProfile handler and decides whether user is logged in or not
//If not, it redirects to login page, else - to the usersProfilePage
func  UserProfileHandler(c *gin.Context) {
	loggedIn := midlwares.CheckLoggedIn(c)
	if loggedIn {
		c.Redirect(http.StatusFound, "/userProfilePage")
		return
	}
	c.Redirect(http.StatusFound, "/login")
}

//ShowUsersProfilePage is a handler function that renders static userProfile page
func ShowUsersProfilePage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": "User's profile page",
			"page": "userprofile",
			"isLoggedIn": isLoggedIn,
		},
	)
}
//ShowBook is a handler function that renders static book page
func ShowBook(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "Book - Description",
		"page" : "book",
		"isLoggedIn": isLoggedIn,

	})
}

//ShowUploadBookPage is a static page hangler func that renders uploadBook page
func ShowUploadBookPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"uploadBookPage.html",
		gin.H{
			"title": "Upload Book Page",
			"page": "uploadpage",
			"isLoggedIn": isLoggedIn,
		},
	)
}

//ShowTakenBooksPage is a static page hangler func that renders taken books page
func ShowTakenBooksPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"takenBooksPage.html",
		gin.H{
			"title": "Taken books",
			"page": "takenBooks",
			"isLoggedIn": isLoggedIn,
		},
	)
}

//ShowCommentsPage is a static page hangler func that renders comments page
func ShowCommentsPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "User comments",
			"page": "comments",
		},
	)
}

//SearchHandler is a static page hangler func that renders search page
func SearchHandler(c *gin.Context)  {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title":      "GoToBooX - search",
		"page":       "search",
		"isLoggedIn": isLoggedIn,
	})
}

//LocationHandler is a static page handler func that renders location page
func LocationHandler(c *gin.Context)  {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title":      "GoToBooX - location",
		"page":       "location",
		"isLoggedIn": isLoggedIn,
	})
}