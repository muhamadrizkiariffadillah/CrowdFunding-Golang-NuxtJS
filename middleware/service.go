package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/authJWT"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

func MiddleWare(authService authJWT.Service, userService users.Service) gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader,"Bearer"){
			response:= helper.APIResponse(http.StatusUnauthorized,"unauthorized","unauthorized action",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
		}

		var tokenString string

		arrayToken := strings.Split(authHeader," ")

		if len(arrayToken) == 2{
			tokenString = arrayToken[0]
		}

		token,err := authService.ValidateToke(tokenString)
		if err != nil {
			response:= helper.APIResponse(http.StatusUnauthorized,"unauthorized","unauthorized action",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
		}

		claim,ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response:= helper.APIResponse(http.StatusUnauthorized,"unauthorized","unauthorized action",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
		}
		
		userId := int(claim["user_id"].(float64))

		user,err := userService.GetUserByID(userId)
		if err != nil {
			response:= helper.APIResponse(http.StatusUnauthorized,"unauthorized","unauthorized action",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
		}

		c.Set("currentUser",user)
	}
	
}
