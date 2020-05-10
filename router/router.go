package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paczulapiotr/goauth2/usecases"
)

// StatusResp response data type
type StatusResp struct {
	Message string `json:"message"`
}

// AuthReq auth request data type
type AuthReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AuthResp auth code response data type
type AuthResp struct {
	Code string `json:"code"`
}

// RunRouter runs service routing
func RunRouter() {
	router := gin.Default()

	router.GET("/status", statusHandler)
	router.POST("/authorize", authorizeHandler)

	router.Run()
	// go runHTTPRedirectRouter("https://localhost:443")
	// router.RunTLS(":443", "cert.pem", "key.pem")
}

func runHTTPRedirectRouter(urlRedirect string) {
	httpRouter := gin.Default()
	httpRouter.GET("/*path", func(c *gin.Context) {
		path := c.Param("path")
		c.Redirect(301, urlRedirect+path)
	})
	httpRouter.Run(":80")
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		StatusResp{"OK"})
}

func authorizeHandler(c *gin.Context) {
	var authParams AuthReq
	c.BindJSON(&authParams)

	code := usecases.LoginForAuthorizationCode(authParams.Login, authParams.Password)
	resp := AuthResp{code}
	c.JSON(http.StatusOK, resp)
}
