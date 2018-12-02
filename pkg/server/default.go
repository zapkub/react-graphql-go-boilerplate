package server

import "github.com/gin-gonic/gin"

// NewAPIRoutes create new gin http server
func NewAPIRoutes() *gin.Engine {

	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]string{"msg": "OK"})
		})
	}

	return router
}
