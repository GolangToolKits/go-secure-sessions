package gosecuresessions

import "net/http"

// Session Session
type Session interface {
	// GetSession(r *http.Request, name string) Session
	Set(key string, value any)
	Get(key string) any
	Save(r *http.Request, w http.ResponseWriter) error

	//GetSession(name string) Session
	//GetSession(r *http.Request, name string) (Session, error)
	//SetOptions(name string, options *ConfigOptions) error
	//Get(r *http.Request, name string) (Session, bool)
	//New(r *http.Request, name string) (Session, bool)
	//Save(r *http.Request, w http.ResponseWriter, s Session) bool

	//Save(r *http.Request, w http.ResponseWriter) bool
	// Name() string
	// Store() SessionStore
}

// // NewSession NewSession
// func NewSession(name string, options *ConfigOptions) Session {
// 	return nil
// }

// // NewSessionWithConfig NewSession
// func NewSessionWithConfig(secretKey string, name string, path string, domain string, maxAge int) {

// }

// // CreateNewSession CreateNewSession
// func CreateNewSession(r *http.Request, w http.ResponseWriter,
// 	secretKey []byte, name string) Session {
// 	return nil
// }

// // GetSession GetSession
// func GetSession(r *http.Request, name string) Session {
// 	return nil
// }

// go mod init github.com/GolangToolKits/go-secure-sessions
