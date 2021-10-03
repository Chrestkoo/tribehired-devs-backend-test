package v1

import (
	"github.com/chrestkoo/project/tribehired-devs-backend-test/controllers/api"
	"github.com/gin-gonic/gin"
)

// Api func
func Api(route *gin.RouterGroup) {
	route.GET("/top-posts-listing/get", api.GetTopPostListing)
	route.GET("/posts-listing/search", api.SearchPostListing)
}
