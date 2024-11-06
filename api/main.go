package main

import (
	"strconv"

	"github.com/ZayenJS/data"
	"github.com/ZayenJS/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		print("Error loading .env file")
	}

	data.SetupDb()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// TODO: refactor + finish this
	r.GET("/food/all", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		perPage := c.DefaultQuery("per-page", "10")

		stmt, err := data.Db.Prepare("SELECT * FROM food LIMIT ? OFFSET ?")

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		convertedPage, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		convertedPerPage, err := strconv.Atoi(perPage)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		rows, err := stmt.Query(perPage, (convertedPage-1)*convertedPerPage)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		foods := []models.Food{}

		for rows.Next() {
			var food models.Food
			err := rows.Scan(&food.ID, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.CreatedAt, &food.UpdatedAt)

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			foods = append(foods, food)
		}

		c.JSON(200, gin.H{"foods": foods})
	})

	r.GET("/food/:id", func(c *gin.Context) {
		var food models.Food
		err := data.Db.QueryRow("SELECT * FROM food WHERE id = ?", c.Param("id")).Scan(&food.ID, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.CreatedAt, &food.UpdatedAt)

		if err != nil {
			c.JSON(404, gin.H{"error": "Food not found"})
			return
		}

		c.JSON(200, gin.H{"food": food})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
