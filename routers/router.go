package routers

import (
	"Gin/controller"
	"strings"

	"Gin/middleware"

	"github.com/gin-gonic/gin"
)

//Route is the information for every URI.
type Route struct {
	Name        string //name of this Riute.
	Method      string //string for the http method. ex) GET,POST, etc...
	Pattern     string //pattern of the URI.
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var BaseController = new(controller.BaseController)
var TestController = new(controller.TestController)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("api/heartbeat", BaseController.HandleHeartbeat)

	v1Group := router.Group("/api/v1").Use(middleware.Authenticate())
	AddRouters(v1Group, v1Routers)

	return router
}

func AddRouters(group gin.IRoutes, routes Routes) {

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}
}
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

var v1Routers = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},
	{
		"HandleGetUserInfo",
		strings.ToUpper("Get"),
		"/user-info",
		TestController.HandleGetUserInfo,
	},
	{
		"HandleAddUserInfo",
		strings.ToUpper("Post"),
		"/user-info",
		TestController.HandleAddUserInfo,
	},
}
