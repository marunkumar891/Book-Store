package middlewares

import (
	"book-store/internal/conf"
	"book-store/internal/errors"
	b64 "encoding/base64"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware(appConf *conf.AppConfig) gin.HandlerFunc {

	return func(c *gin.Context) {

		err := checkIfBasicAuthorized(c, appConf)
		if errors.HasError(&err) {
			c.JSON(err.Status, err)
			c.Abort()
			return
		}
		c.Next()

	}
}

func checkIfBasicAuthorized(c *gin.Context, appConf *conf.AppConfig) errors.RestAPIError {

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return errors.NewUnAuthorizedError("user must be signed in to access this resource")
	}

	basicAuth := strings.Split(authHeader, " ")

	if len(basicAuth) != 2 || basicAuth[1] == "" {
		return errors.NewUnAuthorizedError("Invalid token passed")
	}

	authToken := basicAuth[1]
	uDec, parseError := b64.URLEncoding.DecodeString(authToken)

	if parseError != nil {
		log.Println("Error while Parsing the token: ", parseError.Error())
		return errors.NewUnAuthorizedError("Error while Parsing the token: " + parseError.Error())
	}

	username := strings.Split(string(uDec), ":")[0]
	c.Set("username", username)
	password := strings.Split(string(uDec), ":")[1]

	if (appConf.Server.ADMIN_KEY == username) && (appConf.Server.ADMIN_SECRET == password) {
		return errors.NO_ERROR()
	} else {

		return errors.NewUnAuthorizedError("Invalid clientID and secret")
	}
}
