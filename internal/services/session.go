package services

import "github.com/gorilla/sessions"

// Global store session
var Store = sessions.NewCookieStore([]byte("super-secret-key"))
