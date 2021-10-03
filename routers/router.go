package routers

import (
	"github.com/gin-gonic/gin"
	// "github.com/chrestkoo/project/tribehired-devs-backend-test/middleware/api"
	// "github.com/chrestkoo/project/tribehired-devs-backend-test/middleware/cors"
	v1 "github.com/chrestkoo/project/tribehired-devs-backend-test/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// v1 app api
	apiAppv1 := r.Group("/api/v1")
	{
		// apiAppv1.Use(api.LogAppApiLog())
		v1.Api(apiAppv1) // member api
	}

	return r
}
