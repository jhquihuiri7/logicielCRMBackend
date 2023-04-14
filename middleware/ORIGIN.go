package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ORIGIN() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("Origin"))
		if c.Request.Header.Get("Origin") == "http://localhost:63342" ||
			c.Request.Header.Get("Origin") == "" ||
			c.Request.Header.Get("Origin") == "http://localhost:3000" ||
			c.Request.Header.Get("Origin") == "https://logiciel-crm.netlify.app" {
			if c.Request.Method == "OPTIONS" {
				c.Set("Origin", c.Request.Header.Get("Origin"))
				c.Writer.WriteHeader(http.StatusOK)
				return
			}
			c.Next()
		} else {
			c.Writer.WriteHeader(http.StatusForbidden)
			panic("Origin Forbbiden")
		}

	}
}
