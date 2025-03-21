package controllers

import (
	"github.com/ZayenJS/appHttp"
	"github.com/ZayenJS/repository"
	"github.com/gin-gonic/gin"
)

func GetAllBrands(c *gin.Context) {
	httpResponse := appHttp.NewResponse(c)
	brandRepository := repository.NewBrandRepository()
	brands, err := brandRepository.GetAll()

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	httpResponse.JSON(200, brands)
}
