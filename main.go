package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"notes/admin"
	"notes/database"
	"notes/middlewares"
)

var port string
var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port = "8080"
	DB = database.Connect("admin", "root", "localhost", "dev").DB
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/asset", "./resources")

	router.Use(middlewares.CSRF())
	router.Use(middlewares.CsrfToken())

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":     "Elliot",
			"csrfField": csrf.TemplateField(c.Request),
		})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"csrfField": csrf.TemplateField(c.Request),
		})
	})

	router.POST("/register", func(c *gin.Context) {
		var user admin.User
		c.Bind(&user)
		id := DB.Create(user)
		c.JSON(http.StatusOK, id)
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"csrfField": csrf.TemplateField(c.Request),
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var user admin.User

		account := c.PostForm("account")
		password := c.PostForm("password")

		result := DB.Where("account=?", account).Where("password=?", password).First(&user)

		if result.RowsAffected != 0 {
			c.Redirect(http.StatusFound, "/")
			//c.JSON(http.StatusOK, user)
		} else {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{
				"errorMessage": "Invalid email or password",
			})
			c.JSON(http.StatusOK, "Login Fail")
		}
	})

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
