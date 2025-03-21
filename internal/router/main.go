package router

import "github.com/gin-gonic/gin"

// setRoutes will create all routes
// every group of routes can be defined in its own file
func setRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	addPingRoutes(v1) // for testing
}

// Run Start the API Server
func Run(addr string) {
	r := gin.Default()
	setRoutes(r)

	r.Run(addr)
}
