package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	userID := c.MustGet("user_id")
	role := c.MustGet("role")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user_id": userID,
		"role":    role,
	})
}
