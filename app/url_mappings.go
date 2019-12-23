package app

import (
	"boockstore-user-api/controllers/ping"
	"boockstore-user-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)

}
