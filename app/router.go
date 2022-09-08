package app

import (
	"github.com/hendra24/jwt-template/controller"

	"github.com/hendra24/jwt-template/middlewares"
)

func route() {

	//router.GET("/", controller.Index)
	router.POST("/signup", controller.CreateUser)
	router.GET("/todo", middlewares.TokenAuthMiddleware(), controller.Index)
	router.POST("/login", controller.Login)
	router.PUT("/user", middlewares.TokenAuthMiddleware(), controller.UpdateUser)
	router.POST("/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}
