package gosecuresessions

import (
	"net/http"

	cok "github.com/GolangToolKits/go-secure-cookies"
)

// CookieSession CookieSession
type CookieSession struct {
	id   string
	name string
	// store   Store
	values  map[any]any
	cookies cok.Cookies
	// secretKey string
	path   string
	domain string
	maxAge int
}

// Set Set
func (s *CookieSession) Set(key string, value any) {
	if s.values == nil {
		s.values = make(map[any]any)
	}
	s.values[key] = value
}

// Get Get
func (s *CookieSession) Get(key string) any {
	var rtn any
	if s.values == nil {
		s.values = make(map[any]any)
	} else {
		rtn = s.values[key]
	}
	return rtn
}

// Save Save
func (s *CookieSession) Save(r *http.Request, w http.ResponseWriter) error {

	return nil
}

// // Save Save
// func (s *SecureSession) Save(r *http.Request, w http.ResponseWriter) bool {
// 	return false
// }

// // Name Name
// func (s *SecureSession) Name() string {
// 	return ""
// }

// // Store Store
// func (s *SecureSession) Store() SessionStore {
// 	return nil
// }
