package gosecuresessions

import (
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
				id:      tt.fields.id,
				name:    tt.fields.name,
				values:  tt.fields.values,
				cookies: tt.fields.cookies,
				path:    tt.fields.path,
				domain:  tt.fields.domain,
				maxAge:  tt.fields.maxAge,
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
				key:"test1",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CookieSession{
				id:      tt.fields.id,
				name:    tt.fields.name,
				values:  tt.fields.values,
				cookies: tt.fields.cookies,
				path:    tt.fields.path,
				domain:  tt.fields.domain,
				maxAge:  tt.fields.maxAge,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CookieSession.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
