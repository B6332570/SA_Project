package main

import (
	"github.com/B6332570/SA_Project/controller"
	"github.com/B6332570/SA_Project/entity"

	// "github.com/B6332570/SA_Project/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	// r.Use(CORSMiddleware())

	api := r.Group("/")
	{
		// api.Use(middlewares.Authorizes())
		{
			// User Routes
			api.GET("/users", controller.ListUsers)
			api.GET("/user/:id", controller.GetUser)
			api.PATCH("/users", controller.UpdateUser)
			api.DELETE("/users/:id", controller.DeleteUser)

			// Computer_os Routes
			api.GET("/computer_oss", controller.ListComputer_oss)
			api.GET("/computer_os/:id", controller.GetComputer_os)
			api.POST("/computer_oss", controller.CreateComputer_os)
			api.PATCH("/computer_oss", controller.UpdateComputer_os)
			api.DELETE("/computer_oss/:id", controller.DeleteComputer_os)

			// Computer_reservation Routes
			api.GET("/computer_reservations", controller.ListComputer_reservations)
			api.GET("/computer_reservation/:id", controller.GetComputer_reservation)
			// api.GET("/playlist/watched/user/:id", controller.GetPlaylistWatchedByUser)
			api.POST("/computer_reservations", controller.CreateComputer_reservation)
			api.PATCH("/computer_reservations", controller.UpdateComputer_reservation)
			api.DELETE("/computer_reservations/:id", controller.DeleteComputer_reservation)

			// Computer Routes
			api.GET("/computers", controller.ListComputers)
			api.GET("/computer/:id", controller.GetComputer)
			api.POST("/computers", controller.CreateComputer)
			api.PATCH("/computers", controller.UpdateComputer)
			api.DELETE("/computers/:id", controller.DeleteComputer)

			// Time_com Routes
			api.GET("/time_coms", controller.ListTime_coms)
			api.GET("/time_com/:id", controller.GetTime_com)
			api.POST("/time_coms", controller.CreateTime_com)
			api.PATCH("/time_coms", controller.UpdateTime_com)
			api.DELETE("/time_coms/:id", controller.DeleteTime_com)

		}
	}

	// Signup User Route
	r.POST("/signup", controller.CreateUser)
	// login User Route
	// r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
