package controllers

import (
	"net/http"
	"secure-nas/internal/db"
	"secure-nas/internal/models"
	"secure-nas/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	err := db.DB.QueryRow(
		"SELECT id, username, password_hash, role FROM users WHERE username=?",
		username,
	).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role)

	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	// ✅ Gunakan services.Store global
	session, _ := services.Store.Get(c.Request, "nas_session")
	session.Values["user_id"] = user.ID
	session.Values["role"] = user.Role
	session.Options = &sessions.Options{ // ⚠️ pastikan import github.com/gorilla/sessions
		Path:     "/",
		MaxAge:   60 * 60 * 24, // 1 hari
		HttpOnly: true,
		Secure:   false,
	}
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context) {
	session, _ := services.Store.Get(c.Request, "nas_session")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusFound, "/login")
}
