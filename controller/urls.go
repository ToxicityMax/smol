package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ToxicityMax/smol/models"
	"github.com/ToxicityMax/smol/helpers"
)

// GetallUrls get all urls
func GetallUrls(ctx *gin.Context) {
	var urls []models.Url
	models.DB.Find(&urls)
	ctx.JSON(http.StatusOK, gin.H{"data": urls})
}

// GenShortUrl create shortned url
func GenShortUrl(ctx *gin.Context) {
	// validate input
	var input models.Url
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// todo : create a random number string generator instead of using cripto/rand

	// create url object and save in db
	url := models.Url{LongUrl: input.LongUrl, ShortUrl:helpers.Create_unique_string(3)}
	models.DB.Create(&url)
	ctx.JSON(http.StatusOK, gin.H{"data": url})
}



func Redirect(ctx *gin.Context) {
	var url models.Url
	if err := models.DB.Where("short_url = ?", ctx.Param("slug")).First(&url).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Url"})
		return
	}
	//ctx.Redirect(http.StatusMovedPermanently, url.LongUrl)
	ctx.Redirect(http.StatusMovedPermanently, "https://google.com")

}
