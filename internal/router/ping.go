package router

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/gin-gonic/gin"
)

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		out, _ := exec.Command("ping", "brawlstars.com", "-c 5").Output()

		ps, err := parsePingOutput(string(out))
		if err != nil {
			fmt.Println(err)
		}

		db := database.Db()
		db.Create(ps)

		c.JSON(http.StatusOK, ps)
	})
}

func parsePingOutput(output string) ([]entity.Ping, error) {
	var responses []entity.Ping
	lines := strings.Split(output, "\n")

	var host, ip string

	for _, line := range lines {
		if strings.HasPrefix(line, "PING") {
			parts := strings.Split(line, " ")
			if len(parts) >= 3 {
				host = parts[1]
				ip = strings.Trim(parts[2], "()")
			}
		} else if strings.Contains(line, "time=") {
			parts := strings.Fields(line)
			for _, part := range parts {
				if strings.HasPrefix(part, "time=") {
					timeStr := strings.TrimPrefix(part, "time=")
					time, err := strconv.ParseFloat(strings.TrimSuffix(timeStr, " ms"), 64)
					if err == nil {
						responses = append(responses, entity.Ping{Host: host, IP: ip, Time: time})
					}
				}
			}
		}
	}

	return responses, nil
}
