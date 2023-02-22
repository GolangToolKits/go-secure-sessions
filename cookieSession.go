package gosecuresessions

import (
	"errors"
	"net/http"
	//cok "github.com/GolangToolKits/go-secure-cookies"
)

// CookieSession CookieSession
type CookieSession struct {
	id      string
	name    string
	values  map[any]any
	manager *Manager
	path    string
	domain  string
	maxAge  int
}

// Set Set values
func (s *CookieSession) Set(key string, value any) {
	if s.values == nil {
		s.values = make(map[any]any)
	}
	s.values[key] = value
}

// Get Get values
func (s *CookieSession) Get(key string) any {
	var rtn any
	if s.values == nil {
		s.values = make(map[any]any)
	} else {
		rtn = s.values[key]
	}
	return rtn
}

// Save Save session
func (s *CookieSession) Save(w http.ResponseWriter) error {
	var rtnErr = errors.New("Warning: Failed to save session")
	suc := s.manager.saveSession(w, s)
	if suc {
		rtnErr = nil
	}
	return rtnErr
}
