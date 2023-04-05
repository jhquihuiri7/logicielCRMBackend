package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"logicielcrm/middleware"
	"logicielcrm/models/auth"
	"logicielcrm/models/request"
	"net/http"
	"os"
)

var (
	db     *sql.DB
	err    error
	router *gin.Engine
)

func init() {
	auth.InitKeys()
	db, err = sql.Open("sqlite3", "./data/clients.db")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	router = gin.New()
	router.Use(middleware.CORSMiddleware())
	//router.Use(csrf.Middleware(csrf.Options{
	//	Secret: "secret123",
	//	ErrorFunc: func(c *gin.Context) {
	//		c.String(400, "CSRF token mismatch")
	//		c.Abort()
	//	},
	//}))
	router.POST("/api/login", Login)
	router.GET("/api/validateToken", ValidateToken)
	router.POST("/api/sendBulkMail", BulkMail)

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
func BulkMail(c *gin.Context) {
	var response request.RequestResponse
	resp, err := http.Post(
		"https://mailservicebackend.uc.r.appspot.com/api/bulkMail",
		"application/json",
		c.Request.Body,
	)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.ParseResponse(resp)
	}
	c.Writer.WriteString(response.Marshal())
}
