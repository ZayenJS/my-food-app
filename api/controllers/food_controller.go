package controllers

import (
	"net/http"
	"strconv"
	"strings"

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

	createFoodDto.Name = c.PostForm("name")
	createFoodDto.Brand = c.PostForm("brand")
	createFoodDto.Calories, _ = strconv.ParseFloat(c.PostForm("calories"), 64)
	createFoodDto.Protein, _ = strconv.ParseFloat(c.PostForm("protein"), 64)
	createFoodDto.Carbs, _ = strconv.ParseFloat(c.PostForm("carbs"), 64)
	createFoodDto.Sugar, _ = strconv.ParseFloat(c.PostForm("sugar"), 64)
	createFoodDto.Fat, _ = strconv.ParseFloat(c.PostForm("fat"), 64)
	createFoodDto.SaturatedFat, _ = strconv.ParseFloat(c.PostForm("saturated_fat"), 64)
	createFoodDto.Fiber, _ = strconv.ParseFloat(c.PostForm("fiber"), 64)
	createFoodDto.Sodium, _ = strconv.ParseFloat(c.PostForm("sodium"), 64)

	createFoodDtoErrors := createFoodDto.GetErrors()

	if len(createFoodDtoErrors) > 0 {
		data := make(map[string]interface{})
		data["errors"] = createFoodDtoErrors
		httpResponse.JSON(http.StatusBadRequest, data)
		return
	}

	uuid, err := utils.GenerateUUIDv4()

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	imagePath, err := httpResponse.SaveUploadedFile("image", "./public/uploads/foods", uuid)

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	createFoodDto.NormalizeNames()
	createFoodDto.ImageUrl = imagePath

	brandName := createFoodDto.Brand
	brandRepository := repository.NewBrandRepository()
	brand, err := brandRepository.GetBrandByName(brandName)

	if brand.BrandId == 0 {
		brand := &models.Brand{
			Name: brandName,
		}

		_, err = brandRepository.Create(brand)

		if err != nil {
			httpResponse.Error(500, err)
			return
		}
	}

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	foodRepository := repository.NewFoodRepository()
	food, _ := foodRepository.GetFoodByNameAndBrand(createFoodDto.Name, brand.Name)

	if food != nil {
		data := make(map[string]interface{})
		data["name"] = "food with the same name already exists"
		httpResponse.JSON(400, data)
		return
	}

	// TODO: handle image
	food, err = foodRepository.Create(&createFoodDto, brand.BrandId)

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

	data := make(map[string]interface{})
	data[strings.ToLower(string(food.Name[0]))] = food

	httpResponse.JSON(201, data)
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

	foodRepository := repository.NewFoodRepository()
	if query != "" {
		foods, err := foodRepository.GetFoodsByName(query, convertedPerPage, (convertedPage-1)*convertedPerPage, sort)

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

	foods, err := foodRepository.GetAllFoods(convertedPerPage, (convertedPage-1)*convertedPerPage, sort)

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
