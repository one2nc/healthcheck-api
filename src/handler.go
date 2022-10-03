package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Service is up!")
}

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, "Endpoint doesn't exist")
}

