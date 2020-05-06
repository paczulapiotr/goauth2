package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID   int32  `json:"id"`
	Type string `json:"itemType"`
}

func Start() {
	router := gin.Default()

	jsonHandler := func(c *gin.Context) {
		c.JSON(http.StatusBadRequest,
			Response{100, "Proper JSON item type"})
	}

	router.GET("/", jsonHandler)

	router.GET("/json", jsonHandler)

	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusBadRequest,
			Response{100, "Proper XML item type"})
	})

	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusBadRequest,
			Response{100, "Proper YAML item type"})
	})

	router.Run()
}
