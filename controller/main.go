package controller

import (
	"github.com/ToxicityMax/smol/helpers"
	"github.com/ToxicityMax/smol/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetallUrls get all urls
func GetallUrls(ctx *gin.Context) {
	var urls []models.URL
	models.DB.Find(&urls)
	ctx.JSON(http.StatusOK, urls)
}

type gen_short_url_serializer struct {
	LongUrl        string    `json:"url" binding:"required"`
	ShortUrl       string    `json:"short_url" default:""`
	Password       string    `json:"password" default:""`
	ExpirationDate time.Time `json:"expiration_date"`
}

// GenShortUrl create shortned url
func GenShortUrl(ctx *gin.Context) {
	// validate input
	var body gen_short_url_serializer
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// todo : create a random number string generator instead of using cripto/rand

	// create url object and save in db
	var expiration_date time.Time
	if (body.ExpirationDate != time.Time{}) {
		expiration_date = body.ExpirationDate
	}
	expiration_date = time.Now()
	expiration_date = expiration_date.AddDate(0, 0, 2)
	short := helpers.Create_unique_string(3)

	url := models.URL{
		LongUrl:        body.LongUrl,
		ShortUrl:       short,
		Password:       body.Password,
		ExpirationDate: expiration_date,
	}
	models.DB.Create(&url)

	// change later to address with domain n stuff
	ctx.JSON(http.StatusOK, ctx.Request.Host+"/"+url.ShortUrl)
}

func Redirect(ctx *gin.Context) {
	var url models.URL
	if err := models.DB.Where("short_url = ?", ctx.Param("slug")).First(&url).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Url"})
		return
	}
	url.Visited += 1
	models.DB.Save(&url)
	ctx.Redirect(http.StatusMovedPermanently, url.LongUrl)
}
