package gosecuresessions

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	cok "github.com/GolangToolKits/go-secure-cookies"
)

func TestManager_serializeSession(t *testing.T) {
	cookies, err := cok.NewCookies("dsdfs6dfs61dssdfsdfdsdsfsdsdllsd")
	if err != nil {
		fmt.Println("cookie err: ", err)
	}

	var cs CookieSession
	cs.maxAge = 3600
	cs.name = "test_session"
	cs.values = make(map[any]any)
	cs.values["test1"] = "test11111"
	cs.values["test2"] = "test22222"
	cs.path = "/"
	//cs.cookies = cookies

	type fields struct {
		cookies cok.Cookies
		config  ConfigOptions
	}
	type args struct {
		s Session
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				cookies: cookies,
			},
			args: args{
				s: &cs,
			},
			wantErr: false,
		},
		{
			name: "test 2",
			fields: fields{
				cookies: cookies,
			},
			args: args{
				s: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				cookies: tt.fields.cookies,
				config:  tt.fields.config,
			}
			got, err := m.serializeSession(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.serializeSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "test 1" && got == "" {
				t.Errorf("Manager.serializeSession() = %v, want %v", got, tt.want)
			}
			got2, err := m.deserializeSession(got)
			if tt.name == "test 1" && err != nil {
				t.Fail()
			}
			if got2 != nil {
				rses := got2.(*CookieSession)
				if rses.name != "test_session" || rses.path != "/" || rses.values["test2"] != "test22222" {
					t.Fail()
				}
			}

		})
	}
}

func TestManager_GetSession(t *testing.T) {
	cookies, err := cok.NewCookies("dsdfs6dfs61dssdfsdfdsdsfsdsdllsd")
	if err != nil {
		fmt.Println("cookie err: ", err)
	}

	tr, _ := http.NewRequest("POST", "/test/test1", nil)
	tw := httptest.NewRecorder()

	var cs CookieSession
	cs.maxAge = 3600
	cs.name = "test_session"
	cs.values = make(map[any]any)
	cs.values["test1"] = "test11111"
	cs.values["test2"] = "test22222"
	cs.path = "/"
	//cs.cookies = cookies

	var tman Manager
	tman.cookies = cookies

	scs, err := tman.serializeSession(&cs)
	if err != nil {
		fmt.Println("serial err: ", err)
	}

	cookie := http.Cookie{
		Name:   "test_session",
		Value:  scs,
		MaxAge: 300,
	}

	cookies.Write(tw, cookie)
	cook := tw.Result().Cookies()
	if len(cook) > 0 {
		tr.AddCookie(cook[0])
	}

	///test 3-----
	tr3, _ := http.NewRequest("POST", "/test/test1", nil)
	tw3 := httptest.NewRecorder()

	cookie3 := http.Cookie{
		Name:   "test_session",
		Value:  "this is a session cookie",
		MaxAge: 300,
	}

	cookies.Write(tw3, cookie3)
	cook3 := tw3.Result().Cookies()
	if len(cook3) > 0 {
		tr3.AddCookie(cook3[0])
	}

	//end test 3

	type fields struct {
		cookies cok.Cookies
		config  ConfigOptions
	}
	type args struct {
		r    *http.Request
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Session
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				cookies: cookies,
			},
			args: args{
				r:    tr,
				name: "test_session",
			},
			wantErr: false,
		},
		{
			name: "test 2",
			fields: fields{
				cookies: cookies,
			},
			args: args{
				r:    tr,
				name: "test_session2",
			},
			wantErr: true,
		},
		{
			name: "test 3",
			fields: fields{
				cookies: cookies,
			},
			args: args{
				r:    tr3,
				name: "test_session",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				cookies: tt.fields.cookies,
				config:  tt.fields.config,
			}
			got, err := m.getSession(tt.args.r, tt.args.name)
			if tt.name == "test 1" && (err != nil) != tt.wantErr {
				t.Errorf("Manager.GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "test 1" && got == nil {
				t.Errorf("Manager.GetSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSessionManager(t *testing.T) {
	type args struct {
		secretKey string
		config    ConfigOptions
	}
	tests := []struct {
		name    string
		args    args
		want    SessionManager
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				secretKey: "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd",
				config: ConfigOptions{
					path:   "/",
					maxAge: 3600,
				},
			},
			wantErr: false,
		},
		{
			name: "test 2",
			args: args{
				secretKey: "dsdfs6",
				config: ConfigOptions{
					path:   "/",
					maxAge: 3600,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSessionManager(tt.args.secretKey, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSessionManager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "test 1" && got == nil {
				t.Errorf("NewSessionManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_NewSession(t *testing.T) {
	cookies, err := cok.NewCookies("dsdfs6dfs61dssdfsdfdsdsfsdsdllsd")
	if err != nil {
		fmt.Println("cookie err: ", err)
	}

	tr, _ := http.NewRequest("POST", "/test/test1", nil)

	///test 3--------

	tr3, _ := http.NewRequest("POST", "/test/test1", nil)
	tw3 := httptest.NewRecorder()

	var cs CookieSession
	cs.maxAge = 3600
	cs.name = "test_session"
	cs.values = make(map[any]any)
	cs.values["test1"] = "test11111"
	cs.values["test2"] = "test22222"
	cs.path = "/"
	//cs.cookies = cookies

	var tman Manager
	tman.cookies = cookies

	scs, err := tman.serializeSession(&cs)
	if err != nil {
		fmt.Println("serial err: ", err)
	}

	cookie := http.Cookie{
		Name:   "test_session",
		Value:  scs,
		MaxAge: 300,
	}

	cookies.Write(tw3, cookie)
	cook := tw3.Result().Cookies()
	if len(cook) > 0 {
		tr3.AddCookie(cook[0])
	}

	///end test 3

	type fields struct {
		cookies cok.Cookies
		config  ConfigOptions
	}
	type args struct {
		r    *http.Request
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Session
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				cookies: cookies,
				config: ConfigOptions{
					path:   "/",
					maxAge: 3600,
				},
			},
			args: args{
				r:    tr,
				name: "sess",
			},
		},
		{
			name: "test 2",
			fields: fields{
				cookies: cookies,
				config: ConfigOptions{
					path:   "",
					maxAge: 3600,
				},
			},
			args: args{
				r:    tr,
				name: "sess",
			},
		},
		{
			name: "test 3",
			fields: fields{
				cookies: cookies,
				config: ConfigOptions{
					path:   "/",
					maxAge: 3600,
				},
			},
			args: args{
				r:    tr3,
				name: "test_session",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				cookies: tt.fields.cookies,
				config:  tt.fields.config,
			}
			if got := m.NewSession(tt.args.r, tt.args.name); got == nil {
				t.Errorf("Manager.NewSession() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestManager_saveSession(t *testing.T) {
	cookies, err := cok.NewCookies("dsdfs6dfs61dssdfsdfdsdsfsdsdllsd")
	tw := httptest.NewRecorder()
	tr, _ := http.NewRequest("POST", "/test/test1", nil)

	type TestObj struct {
		ID   int64
		Name string
		Age  int64
	}
	var tobj TestObj
	tobj.ID = 21
	tobj.Name = "Hacker"
	tobj.Age = 16
	gob.Register(TestObj{})

	var cs CookieSession
	cs.maxAge = 3600
	cs.name = "test_session"
	cs.values = make(map[any]any)
	cs.values["test1"] = "test11111"
	cs.values["test2"] = "test22222"
	cs.values["test3"] = tobj
	cs.path = "/"

	if err != nil {
		fmt.Println("cookie err: ", err)
	}

	type fields struct {
		cookies cok.Cookies
		config  ConfigOptions
	}
	type args struct {
		w http.ResponseWriter
		s Session
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				cookies: cookies,
			},
			args: args{
				w: tw,
				s: &cs,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				cookies: tt.fields.cookies,
				config:  tt.fields.config,
			}
			if got := m.saveSession(tt.args.w, tt.args.s); got != tt.want {
				t.Errorf("Manager.saveSession() = %v, want %v", got, tt.want)
			}

			//cookies.Write(tw, cookie)
			cook := tw.Result().Cookies()
			if len(cook) > 0 {
				tr.AddCookie(cook[0])
			}
			ses := m.NewSession(tr, "test_session")
			if ses.Get("test2") != "test22222" {
				t.Fail()
			}

			obj := ses.Get("test3")
			ob := obj.(TestObj)
			if ob.ID != 21 || ob.Name != "Hacker" || ob.Age != 16 {
				t.Fail()
			}
		})
	}
}
