package controllers

import (
	"net/http"
	"strconv"

	"github.com/ZayenJS/appHttp"
	"github.com/ZayenJS/models"
	"github.com/ZayenJS/utils"
	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {
	httpResponse := appHttp.NewResponse(c)
	/**
	  TODO: Implement the CreateFood function
		1. Create the CreateFootDto struct
		2. Bind the request body to the CreateFootDto struct
		3. Validate the CreateFootDto struct
		4. Create a new Food struct
		5. Assign the values from the CreateFootDto struct to the Food struct
		6. Instantiate Food Repository
		7. Call the CreateFood method from the Food Repository
		8. Handle the error
		9. Return a 201 status code
	*/

	// go routine example
	// dataCh := make(chan bool)

	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	fmt.Println("Start data generation")
	// 	defer wg.Done()
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)

	// 		dataCh <- true
	// 	}

	// 	close(dataCh)
	// }()

	// go func() {
	// 	fmt.Println("Start data processing")
	// 	for range dataCh {
	// 		println("Data received")
	// 	}
	// }()

	// wg.Wait()

	httpResponse.JSON(201, gin.H{"test": 1})
}

func GetAllFoods(c *gin.Context) {
	httpResponse := appHttp.NewResponse(c)

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per-page", "10")
	sort := utils.ValidateSqlSortType(c.DefaultQuery("sort", "asc"))
	indexBy := c.Query("index-by")

	convertedPage, err := strconv.Atoi(page)
	if err != nil {
		httpResponse.Error(500, err)
		return
	}
	convertedPerPage, err := strconv.Atoi(perPage)
	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	query := c.Query("name")

	if query != "" {
		foods, err := models.GetFoodsByName(query, convertedPerPage, (convertedPage-1)*convertedPerPage, sort)

		if err != nil {
			httpResponse.Error(500, err)
			return
		}

		if indexBy == "letter" {
			httpResponse.JSON(http.StatusOK, utils.IndexFoodsByLetter(foods))
			return
		}

		httpResponse.JSON(200, foods)
		return
	}

	foods, err := models.GetAllFoods(convertedPerPage, (convertedPage-1)*convertedPerPage, sort)

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	if indexBy == "letter" {
		httpResponse.JSON(http.StatusOK, utils.IndexFoodsByLetter(foods))
		return
	}

	httpResponse.JSON(200, foods)
}
