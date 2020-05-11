package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paczulapiotr/goauth2/usecases"
)

// RunRouter runs service routing
func RunRouter() {
	router := gin.Default()

	router.GET("/status", statusHandler)
	router.POST("/authorize", authorizeHandler)
	router.POST("/register", registerHandler)
	router.POST("/code", useCodeHandler)
	router.POST("/revoke", revokeHandler)
	router.POST("/refresh", refreshHandler)
	router.POST("/check", checkTokenHandler)

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

func checkTokenHandler(c *gin.Context) {
	var payload CheckAccessTokenReq
	c.BindJSON(&payload)

	err := usecases.CheckAccessToken(payload.AccessToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func revokeHandler(c *gin.Context) {
	var payload RefreshTokenReq
	c.BindJSON(&payload)

	err := usecases.RevokeRefreshToken(payload.RefreshToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func refreshHandler(c *gin.Context) {
	var payload RefreshTokenReq
	c.BindJSON(&payload)

	accessToken, err := usecases.RefreshAccessToken(payload.RefreshToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		response := RefreshAccessTokenResp{accessToken}
		c.JSON(http.StatusOK, response)
	}
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		StatusResp{"OK"})
}

func authorizeHandler(c *gin.Context) {
	var authParams AuthReq
	c.BindJSON(&authParams)

	code, err := usecases.LoginForAuthorizationCode(authParams.Login, authParams.Password)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
	} else {
		resp := AuthResp{code}
		c.JSON(http.StatusOK, resp)
	}
}

func registerHandler(c *gin.Context) {
	var registerParams RegisterReq
	c.BindJSON(&registerParams)
	err := usecases.RegisterUser(registerParams.Login, registerParams.Password)

	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func useCodeHandler(c *gin.Context) {
	var requestData UseCodeReq
	c.BindJSON(&requestData)

	accessToken,
		validUntil,
		refreshToken,
		refreshValidUntil,
		err := usecases.UseAuthorizationCode(requestData.Code)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		response := UseCodeResp{
			AccessToken:       accessToken,
			ValidUntil:        validUntil,
			RefreshToken:      refreshToken,
			RefreshValidUntil: refreshValidUntil,
		}
		c.JSON(http.StatusOK, response)
	}
}
