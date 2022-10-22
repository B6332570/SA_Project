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
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		// router.Use(middlewares.Authorizes())
		{
			// User Routes
			router.GET("/users", controller.ListUsers)
			router.GET("/user/:id", controller.GetUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			// Computer_os Routes
			router.GET("/computer_oss", controller.ListComputer_oss)
			router.GET("/computer_os/:id", controller.GetComputer_os)
			router.POST("/computer_oss", controller.CreateComputer_os)
			router.PATCH("/computer_oss", controller.UpdateComputer_os)
			router.DELETE("/computer_oss/:id", controller.DeleteComputer_os)

			// Computer_reservation Routes
			router.GET("/computer_reservations", controller.ListComputer_reservations)
			router.GET("/computer_reservation/:id", controller.GetComputer_reservation)
			// router.GET("/playlist/watched/user/:id", controller.GetPlaylistWatchedByUser)
			router.POST("/computer_reservation", controller.CreateComputer_reservation)
			router.PATCH("/computer_reservations", controller.UpdateComputer_reservation)
			router.DELETE("/computer_reservations/:id", controller.DeleteComputer_reservation)

			// Computer Routes
			router.GET("/computers", controller.ListComputers)
			router.GET("/computer/:id", controller.GetComputer)
			router.POST("/computers", controller.CreateComputer)
			router.PATCH("/computers", controller.UpdateComputer)
			router.DELETE("/computers/:id", controller.DeleteComputer)

			// Time_com Routes
			router.GET("/time_coms", controller.ListTime_coms)
			router.GET("/time_com/:id", controller.GetTime_com)
			router.POST("/time_coms", controller.CreateTime_com)
			router.PATCH("/time_coms", controller.UpdateTime_com)
			router.DELETE("/time_coms/:id", controller.DeleteTime_com)

		}
	}

	// // Signup User Route
	// r.POST("/signup", controller.CreateUser)
	// // // login User Route
	// r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("0.0.0.0:8080")
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
