package healthleft

import (
	"cryptchat/internal/health/core"
	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {

	health, err := healthcore.GetHealth()
	if err != nil {
		c.JSON(500, health)
		return
	}
	c.JSON(200, health)
}
