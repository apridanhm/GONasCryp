package middlewares

import (
	"net/http"
	"secure-nas/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthRequired(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := services.Store.Get(c.Request, "nas_session")
		userRole, ok := session.Values["role"].(string)
		if !ok || userRole == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		if role != "" && userRole != role {
			c.String(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}
		c.Set("user_id", session.Values["user_id"])
		c.Set("role", userRole)
		c.Next()
	}
}
