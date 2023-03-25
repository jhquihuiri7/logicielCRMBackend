package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"logicielcrm/models/auth"
	"net/http"
	"os"
	"time"
)

var (
	db  *sql.DB
	err error
)

func init() {
	auth.InitKeys()
	db, err = sql.Open("sqlite3", "./data/clients.db")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Content-Length", "Content-type"},
		ExposeHeaders:    []string{"Content-Length", "Content-type"},
		AllowCredentials: true,

		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("/api/login", Login)
	router.GET("/api/validateToken", ValidateToken)

	port := os.Getenv("PORT")
	if err = http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
func Login(c *gin.Context) {
	var request auth.AuthRequest
	request.ParseAuth(c.Request.Body)
	response := request.Login(db)
	c.Writer.WriteString(response.Marshal())
}
func ValidateToken(c *gin.Context) {
	var request auth.AuthResponse
	response := request.ValidateToken(c.Request, db)
	c.Writer.WriteString(response.Marshal())
}
