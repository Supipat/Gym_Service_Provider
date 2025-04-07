package main

import (
	"gymservice/connectdb"
	"gymservice/handlers"
	"gymservice/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	connectdb.ConnectDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/register", handlers.RegisterClient)
	r.POST("/login", handlers.CreateLogin)

	api := r.Group("/api")
	{
		api.GET("/gymclass", func(c *gin.Context) {
			params := models.QueryParams{
				Search: c.Query("search"),
				Sort:   c.Query("sort"),
				Order:  c.Query("order"),
			}

			if gymIDStr := c.Query("gym_id"); gymIDStr != "" {
				if gymID, err := strconv.Atoi(gymIDStr); err == nil {
					params.GymID = gymID
				}
			}

			handlers.GetAllGym(c, params)
		})
		api.GET("/gym/:id", func(c *gin.Context) {
			gymId := c.Param("id")

			gym_id, err := strconv.Atoi(gymId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
				return
			}

			handlers.GetGymId(c, gym_id)
		})
		api.GET("/affiliator", handlers.GetClient)
		api.GET("/affiliator/:email", handlers.GetAffiliatorByEmail)
		api.GET("/log", handlers.GetLog)
		api.POST("/gym", handlers.CreateGym)
		api.POST("/gymclass", handlers.CreateGymClass)
		api.POST("/clicklog", handlers.LogClick)
		api.GET("/clicklog", handlers.GetClickLogs)
	}

	auth := r.Group("/")
	auth.Use(handlers.AuthMiddleware(), handlers.LogMiddleware())
	{
		auth.GET("/gymclass/:id", func(c *gin.Context) {
			params := models.QueryParams{
				Search: c.Query("search"),
				Sort:   c.Query("sort"),
				Order:  c.Query("order"),
			}

			gymId := c.Param("id")

			gym_id, err := strconv.Atoi(gymId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
				return
			}
			handlers.GetGymClass(c, gym_id, params)
		})
		auth.GET("/gymclass/:id/:idc", func(c *gin.Context) {
			gymId := c.Param("id")
			classId := c.Param("idc")

			gym_id, err := strconv.Atoi(gymId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
				return
			}

			class_id, err := strconv.Atoi(classId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
				return
			}
			handlers.GetGymClassId(c, gym_id, class_id)
		})
	}

	r.Run(":8088")
}
