package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

type UserService struct {
	UsersRepo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		UsersRepo: repository,
	}
}

//UserGetHandler gets users data from database using unique email that is stored in cookie
//if there is no email in coolie that means that session is over
func (s *UserService) UserGetHandler(c *gin.Context) {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	email := convertEmailString(emailCookie.Value)
	user, err := s.UsersRepo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

//UserDeleteHandler deletes user from database. Uses DELETE method.
func (s *UserService) UserDeleteHandler(c *gin.Context) {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	email := convertEmailString(emailCookie.Value)
	if err := s.UsersRepo.DeleteUserByEmail(email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.SetCookie("email", "", -1, "", "", false, true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

/* UserUpdateHandler is a handler function that updates users info in database. Uses PUT method
Input example for update
{
	"id": 1,
	"nickname": "Denchick",
	"email": "away4ppel@den.ua",
	"password": "pass",
	"registrDate": "2018-01-01"
}
 */
func (s *UserService) UserUpdateHandler(c *gin.Context) {
	var u repository.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	email := convertEmailString(emailCookie.Value)

	if err := s.UsersRepo.UpdateUserByEmail(u, email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	c.SetCookie("email", u.Email, 15000, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
	return
}

//This function was created because cookies gives '%40' instead of '@' when read the email. It converts
func convertEmailString(email string) (string) {
	indexOfPercentSymb := strings.IndexRune(email, '%')
	runes := []rune(email)
	runes[indexOfPercentSymb] = '@'
	runes = append(runes[:indexOfPercentSymb+1], runes[indexOfPercentSymb+2:]...) //deletes 4
	runes = append(runes[:indexOfPercentSymb+1], runes[indexOfPercentSymb+2:]...) //deletes 0
	return string(runes)
}
