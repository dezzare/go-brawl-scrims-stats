package router

import (
	"math/rand/v2"
	"net/http"
	"sync"
	"time"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/gin-gonic/gin"
)

// addPingRoutes make nonsense things, just to satisfy curiosity and testing
func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		var wg sync.WaitGroup
		var ps []entity.Ping

		for i := 0; i <= 3; i++ {
			wg.Add(3)
			go since(&ps, &wg)
			go since(&ps, &wg)
			go since(&ps, &wg)
		}

		wg.Wait()

		saveToDB(&ps)

		c.JSON(http.StatusOK, ps)
	})
}

func since(ps *[]entity.Ping, wg *sync.WaitGroup) {
	defer wg.Done()

	d := time.Now()

	for i := 0; i <= 5; i++ {
		s := rand.IntN(10)

		elapse := time.Since(d).Seconds()
		elapse = elapse*1000 + float64(i+1)/10
		time.Sleep(time.Duration(s) * time.Millisecond)
		ping := entity.Ping{
			TimeInMs: elapse,
		}
		*ps = append(*ps, ping)
	}

}

func saveToDB(ps *[]entity.Ping) {
	db := database.Db()
	db.Create(ps)
}
