package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDataRouter(rg *gin.RouterGroup) {
	myData := rg.Group("/edi-vazquez-luna")

	myData.GET("/about-me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"aboutMe": "data.AboutMe2"})
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
