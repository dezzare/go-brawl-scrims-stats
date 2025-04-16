package api

import (
	"os"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
	DB     *service.DB
}

// New set app routes/templates and return router
func New() (*API, error) {

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	a := API{
		Router: gin.Default(),
		DB:     db,
	}

	a.Router.Use(database.Inject(db))

	a.loadTemplates()
	a.setRoutes()

	return &a, nil
}

// Load all HTML templates
func (a *API) loadTemplates() {
	a.Router.LoadHTMLGlob("views/**/*")
	a.Router.Static("/v1/css", "views/css")
}

// setRoutes will create all routes
// every group of routes can be defined in its own file
func (a *API) setRoutes() {
	v1 := a.Router.Group("/v1")
	a.addHomeRoute(v1)
	a.addPlayerRoute(v1)
	a.addTeamRoute(v1)
}

func (a *API) Run() {
	port := os.Getenv("APP_PORT")
	a.Router.Run(":" + port)
}
