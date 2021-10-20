package router

import "github.com/gin-gonic/gin"

type router struct {
	route *gin.RouterGroup
}

type MappingRouter interface {
	GET(url string, f func(c *gin.Context))
	POST(url string, f func(c *gin.Context))
	PUT(url string, f func(c *gin.Context))
	DELETE(url string, f func(c *gin.Context))
}

func NewBaseRouter(r *gin.RouterGroup) MappingRouter {
	return &router{
		route: r,
	}
}

func (r *router) GET(url string, f func(c *gin.Context)) {
	r.route.GET(url, f)
}

func (r *router) POST(url string, f func(c *gin.Context)) {
	r.route.POST(url, f)
}

func (r *router) PUT(url string, f func(c *gin.Context)) {
	r.route.PUT(url, f)
}

func (r *router) DELETE(url string, f func(c *gin.Context)) {
	r.route.DELETE(url, f)
}
