package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	httpCommon "github.com/kevinnaserwan/coursphere-api/internal/http/server"
	errorCommon "github.com/kevinnaserwan/coursphere-api/internal/util/error"
	"github.com/kevinnaserwan/coursphere-api/internal/util/jwt"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]
			// if err can be casted to ClientError, then it is a client error
			if clientError, ok := err.Err.(*errorCommon.ClientError); ok {
				c.JSON(clientError.Code, httpCommon.Error{
					Message: clientError.Message,
				})
			} else if err.IsType(gin.ErrorTypeBind) {
				c.JSON(400, httpCommon.Error{
					Message: err.Err.Error(),
				})
			} else if err.IsType(gin.ErrorTypePrivate) {
				fmt.Println(err.Err.Error())
				c.JSON(500, httpCommon.Error{
					Message: "Internal server error",
				})
			} else {
				fmt.Println(err.Err.Error())
				c.JSON(500, httpCommon.Error{
					Message: "Internal server error",
				})
			}
		}
	}
}

func JWTAuth(j *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(
				errorCommon.NewForbidden("you are not authenticated to access this route"),
			)
			c.Abort()
			return
		} else if len(authHeader) < BEARER {
			c.Error(errorCommon.NewBadRequest("authorization header not valid"))
			c.Abort()
			return
		}

		tokenString := authHeader[BEARER:]
		claims, err := j.VerifyAuthToken(tokenString)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("user_id", claims.Identifier)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

func RoleAuth(roles string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("user_role")

		if !strings.Contains(roles, userRole) {
			c.Error(
				errorCommon.NewUnauthorized("you are not authorized to access this route"),
			)
			c.Abort()
			return
		}

		c.Next()
	}
}
