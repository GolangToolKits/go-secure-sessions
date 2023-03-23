package gosecuresessions

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"net/http"

	cok "github.com/GolangToolKits/go-secure-cookies"
)

const (
	cookieSession int = 0
)

// SessionManager SessionManager
type SessionManager interface {
	NewSession(r *http.Request, name string) Session
}

// Manager Manager
type Manager struct {
	cookies cok.Cookies
	config  ConfigOptions
}

// NewSessionManager securekey must be at least 16 char long
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func NewSessionManager(secretKey string, config ConfigOptions) (SessionManager, error) {
	var rtn SessionManager
	var rtnErr error
	cookies, err := cok.NewCookies(secretKey)
	if err == nil {
		var m Manager
		m.cookies = cookies
		m.config = config
		rtn = &m
	} else {
		rtnErr = err
	}
	return rtn, rtnErr
}

// NewSession NewSession
// Gets session from request if one exists
// Creates a new session if one doesn't exist in request
func (m *Manager) NewSession(r *http.Request, name string) Session {
	var rtn Session
	if m.config.SessionType == cookieSession {
		eses, err := m.getSession(r, name)
		if eses != nil && err == nil {
			rtn = eses
		} else {
			var ses CookieSession
			ses.manager = m
			ses.name = name
			if m.config.Path == "" {
				ses.path = "/"
			} else {
				ses.path = m.config.Path
			}
			ses.maxAge = m.config.MaxAge
			ses.values = make(map[any]any)
			rtn = &ses
		}
	}
	return rtn
}

// SaveSession Saves the Session to a secure cookie
func (m *Manager) saveSession(w http.ResponseWriter, s Session) bool {
	var rtn bool
	if s != nil {
		cs := s.(*CookieSession)
		ss, err := m.serializeSession(s)
		if err == nil {
			var ck http.Cookie
			ck.Name = cs.name
			ck.Value = ss
			ck.Path = cs.path
			ck.MaxAge = cs.maxAge
			ck.Domain = cs.domain
			err := m.cookies.Write(w, ck)
			if err == nil {
				rtn = true
			}
		}
	}
	return rtn
}

// GetSession GetSession
func (m *Manager) getSession(r *http.Request, name string) (Session, error) {
	var rtn Session
	var rtnErr error
	sc, err := m.cookies.Read(r, name)
	if err == nil && sc != "" {
		ses, err := m.deserializeSession(sc)
		if err == nil {
			rtn = ses
		} else {
			rtnErr = err
		}
	} else {
		rtnErr = errors.New("Could not read session cookie named " + name)
	}
	return rtn, rtnErr
}

func (m *Manager) serializeSession(s Session) (string, error) {
	var rtn string
	var rtnErr error
	var esus = false
	if s != nil {
		cs := s.(*CookieSession)
		b := new(bytes.Buffer)
		encoder := gob.NewEncoder(b)
		err := encoder.Encode(cs.id)
		if err == nil {
			err := encoder.Encode(cs.maxAge)
			if err == nil {
				err := encoder.Encode(cs.domain)
				if err == nil {
					err := encoder.Encode(cs.name)
					if err == nil {
						err := encoder.Encode(cs.path)
						if err == nil {
							err := encoder.Encode(cs.values)
							if err == nil {
								esus = true
								rtn = base64.StdEncoding.EncodeToString(b.Bytes())
							}
						}
					}
				}
			}
		}
	}
	if !esus {
		rtnErr = errors.New("Failure: There was a problem in serializing the secure session")
	}
	return rtn, rtnErr
}

func (m *Manager) deserializeSession(ss string) (Session, error) {
	var rtn Session
	var rtnErr error
	var cses CookieSession
	var esus = false

	bs, err := base64.StdEncoding.DecodeString(ss)
	if err == nil {
		b := bytes.NewBuffer(bs)
		decoder := gob.NewDecoder(b)
		err := decoder.Decode(&cses.id)
		if err == nil {
			err := decoder.Decode(&cses.maxAge)
			if err == nil {
				err := decoder.Decode(&cses.domain)
				if err == nil {
					err := decoder.Decode(&cses.name)
					if err == nil {
						err := decoder.Decode(&cses.path)
						if err == nil {
							err := decoder.Decode(&cses.values)
							if err == nil {
								//cookies, err := cok.NewCookies(secretKey)
								//cses.cookies = m.cookies
								cses.manager = m
								esus = true
								rtn = &cses
							}
						}
					}
				}
			}
		}
	}
	if !esus {
		rtnErr = errors.New("Failure: There was a problem in deserializing the secure session")
	}
	return rtn, rtnErr
}
