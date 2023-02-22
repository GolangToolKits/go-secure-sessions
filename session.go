package gosecuresessions

import "net/http"

// Session Session
type Session interface {
	Set(key string, value any)
	Get(key string) any
	Save(w http.ResponseWriter) error
}

// go mod init github.com/GolangToolKits/go-secure-sessions
