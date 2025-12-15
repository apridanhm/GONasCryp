package bootstrap

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func EnsureAdmin(db *sql.DB, username, password string) {
	var count int
	db.QueryRow(
		"SELECT COUNT(*) FROM users WHERE role='admin'",
	).Scan(&count)

	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO users (id, username, password_hash, role)
		VALUES (?, ?, ?, 'admin')
	`,
		uuid.New().String(),
		username,
		string(hash),
	)

	if err != nil {
		log.Fatal("failed create admin:", err)
	}

	log.Println("✅ Admin bootstrap created")
}
