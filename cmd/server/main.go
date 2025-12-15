package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"secure-nas/internal/app"
	"secure-nas/internal/bootstrap"
	"secure-nas/internal/controllers"
	"secure-nas/internal/db"
	"secure-nas/internal/middlewares"
)

func main() {
	cfg := app.LoadConfig()

	// Init DB
	if err := db.Init(cfg.Database.DSN); err != nil {
		log.Fatal(err)
	}

	// Ensure admin
	bootstrap.EnsureAdmin(db.DB, cfg.Security.AdminUsername, cfg.Security.AdminPassword)

	// Default storage path
	ensureStorageSetting(cfg.Storage.DefaultPath)

	// Gin engine
	r := gin.Default()

	// ===== Templates =====
	wd, _ := os.Getwd()
	projectRoot := wd
	authTemplate := filepath.Join(projectRoot, "views", "auth", "login.html")
	dashboardTemplate := filepath.Join(projectRoot, "views", "dashboard", "index.html")
	r.LoadHTMLFiles(authTemplate, dashboardTemplate)
	log.Println("✅ Templates loaded:", authTemplate, dashboardTemplate)
	// =====================

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Auth routes
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.LoginPost)
	r.POST("/logout", controllers.Logout)

	// Dashboard route
	r.GET("/dashboard", middlewares.AuthRequired(""), controllers.Dashboard)

	// Run server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Println("🚀 Secure NAS running on", addr)
	r.Run(addr)
}

func ensureStorageSetting(path string) {
	_, err := db.DB.Exec("INSERT IGNORE INTO settings (k,v) VALUES ('storage_path', ?)", path)
	if err != nil {
		log.Fatal("failed set default storage path:", err)
	}
}
