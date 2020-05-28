package main

type routesTestDefiner struct {
	Name	string
	Port	int
}

var routesTestTemplate = &RoutesTest {
	`package main

	import (
		"net/http"
		"net/http/httptest"
		"strings"
		"testing"
	)
	
	func Test_handleVersion(t *testing.T) {
	
		srv := server{router: http.NewServeMux()}
		srv.routes()
	
		type test struct {
			name       string
			r          *http.Request
			wantStatus int
			wantError  string
		}
	
		tests := []test{
			{
				name: "successful request",
				r: func() *http.Request {
					req, _ := http.NewRequest("GET", "/version", nil)
					return req
				}(),
				wantStatus: http.StatusOK,
				wantError:  "{\"version\":\"v\"}",
			},
		}
	
		for _, tt := range tests {
	
			t.Run(tt.name, func(t *testing.T) {
	
				rr := httptest.NewRecorder()
				srv.router.ServeHTTP(rr, tt.r)
	
				if rr.Code != tt.wantStatus {
					t.Errorf("handler returned wrong status code: got %v, want %v",
						rr.Code, tt.wantStatus)
				}
	
				if !strings.Contains(rr.Body.String(), tt.wantError) {
					t.Errorf("handler returned unexpected body: got %v, should contain %v",
						rr.Body.String(), tt.wantError)
				}
	
			})
		}
	}`,
}