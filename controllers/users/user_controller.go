package users

import (
	"boockstore-user-api/domain/users"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	c.String(http.StatusNotImplemented, "Implement me!")
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
