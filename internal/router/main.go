package router

import "github.com/gin-gonic/gin"

var router = gin.Default()

// Run will start the server
func Run() {
	setRoutes()
	router.Run(":5000")
}

// setRoutes will create all routes
// every group of routes can be defined in its own file
func setRoutes() {
	v1 := router.Group("/v1")
	addPingRoutes(v1) // for testing
}
