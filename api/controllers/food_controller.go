package controllers

import (
	"net/http"
	"strconv"

	"github.com/ZayenJS/appHttp"
	"github.com/ZayenJS/dto"
	"github.com/ZayenJS/models"
	"github.com/ZayenJS/repository"
	"github.com/ZayenJS/utils"
	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {
	httpResponse := appHttp.NewResponse(c)
	var createFoodDto dto.CreateFoodDTO

	err := c.ShouldBindJSON(&createFoodDto)
	if err != nil {
		httpResponse.Error(400, err)
		return
	}

	brandName := createFoodDto.Brand
	brand, err := repository.NewBrandRepository().GetBrandByName(brandName)

	if brand == nil {
		// TODO: create brand if not exist
	}

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	// TODO: handle image and category
	food, err := repository.NewFoodRepository().Create(&createFoodDto, brand.BrandId)

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

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

	httpResponse.JSON(201, food)
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
