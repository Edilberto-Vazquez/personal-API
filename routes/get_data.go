package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/personal-API/schemas"
	"github.com/Edilberto-Vazquez/personal-API/services"
	"github.com/Edilberto-Vazquez/personal-API/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDataRouter(rg *gin.RouterGroup) {
	myData := rg.Group("/my-api")

	myData.GET("/about-me", func(c *gin.Context) {
		var aboutMe schemas.AboutMeSchema
		var data primitive.M
		var err error
		if err = c.ShouldBindQuery(&aboutMe); err != nil {
			utils.ErrorMessage(err, c)
		} else {
			data, err = services.AboutMeFindOne(aboutMe.Alias)
		}
		if err != nil {
			utils.ErrorMessage(err, c)
		} else {
			c.JSON(http.StatusOK, gin.H{"aboutMe": data})
		}
	})

	myData.GET("/resume", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"resume": "data.Resume"})
	})

	myData.GET("/portafolio", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"portafolio": "data.Portafolio",
		})
	})
}
