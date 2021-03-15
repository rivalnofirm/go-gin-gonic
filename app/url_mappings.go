package app

import "github.com/rivalnofirm/bookstore_users-api/controllers/users"

func mapsUrls() {
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.Create)
}
