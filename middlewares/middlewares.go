package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hendra24/jwt-template/auth"
	"github.com/hendra24/jwt-template/model"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		au, err := auth.ExtractTokenAuth(c.Request)
		if err != nil {

			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}

		_, err = model.Model.FetchAuth(au)
		if err != nil {
			log.Println("User with auth uuid " + au.AuthUuid + " NOT ACTIVE")
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		log.Println("User with auth uuid " + au.AuthUuid + " ACTIVE")
		c.Next()
	}
}
