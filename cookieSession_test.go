package gosecuresessions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	cok "github.com/GolangToolKits/go-secure-cookies"
)

func TestCookieSession_Set(t *testing.T) {
	type fields struct {
		id      string
		name    string
		values  map[any]any
		cookies cok.Cookies
		path    string
		domain  string
		maxAge  int
	}
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			fields: fields{},
			args: args{
				key:   "test1",
				value: "a test 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CookieSession{
				id:     tt.fields.id,
				name:   tt.fields.name,
				values: tt.fields.values,
				//cookies: tt.fields.cookies,
				path:   tt.fields.path,
				domain: tt.fields.domain,
				maxAge: tt.fields.maxAge,
			}
			s.Set(tt.args.key, tt.args.value)
			if tt.name == "test 1" && s.Get("test1") != "a test 1" {
				t.Fail()
			}
		})

	}
}

func TestCookieSession_Get(t *testing.T) {
	type fields struct {
		id      string
		name    string
		values  map[any]any
		cookies cok.Cookies
		path    string
		domain  string
		maxAge  int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				key: "test1",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CookieSession{
				id:     tt.fields.id,
				name:   tt.fields.name,
				values: tt.fields.values,
				//cookies: tt.fields.cookies,
				path:   tt.fields.path,
				domain: tt.fields.domain,
				maxAge: tt.fields.maxAge,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CookieSession.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCookieSession_Save(t *testing.T) {
	var cf ConfigOptions
	cf.MaxAge = 3600
	m, err := NewSessionManager("dsdfs6dfs61dssdfsdfdsdsfsdsdllsd", cf)
	if err != nil {
		fmt.Println(err)
	}
	v := make(map[any]any)
	v["test1"] = "test11111"
	v["test2"] = "test22222"

	tw := httptest.NewRecorder()
	tr, _ := http.NewRequest("POST", "/test/test1", nil)

	type fields struct {
		id      string
		name    string
		values  map[any]any
		manager *Manager
		path    string
		domain  string
		maxAge  int
	}
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				name:    "test_session",
				values:  v,
				manager: m.(*Manager),
				path:    "/",
				maxAge:  3600,
			},
			args: args{
				w: tw,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CookieSession{
				id:      tt.fields.id,
				name:    tt.fields.name,
				values:  tt.fields.values,
				manager: tt.fields.manager,
				path:    tt.fields.path,
				domain:  tt.fields.domain,
				maxAge:  tt.fields.maxAge,
			}
			if err := s.Save(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("CookieSession.Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			cook := tw.Result().Cookies()
			if len(cook) > 0 {
				tr.AddCookie(cook[0])
			}
			ses := m.NewSession(tr, "test_session")
			if ses.Get("test2") != "test22222" {
				t.Fail()
			}
		})
	}
}
