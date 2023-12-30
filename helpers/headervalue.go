
package helpers

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAuthorizationHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ""
	}

	return headerParts[1]
}
