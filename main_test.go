package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t * testing.T) {
	mux := setupServer()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", rec.Code)
	}

	expectedBody := "Hello World\n"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, rec.Body.String())
	}

	expectedContentType := "text/plain; charset=utf-8"
	if rec.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected header %s, got %s", expectedContentType, rec.Header().Get("Content-Type"))
	}
}

func TestHomeHandler(t * testing.T) {
	mux := setupServer()

	req := httptest.NewRequest(http.MethodGet, "/home", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", rec.Code)
	}

	expectedBody := "Hello Home\n"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, rec.Body.String())
	}
}

func TestConcurrentClients(t * testing.T) {
	mux := setupServer()

	tests := []struct {
		route			string
		expectedBody 	string
		expectedCode	int
	}{
		{ "/", "Hello World\n", http.StatusOK },
		{ "/home", "Hello Home\n", http.StatusOK },
//		{ "/unknown", "404 Page not found\n", http.StatusNotFound },
	}


	for _, tc := range tests {
		tc := tc
		t.Run(tc.route, func(t * testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, tc.route, nil)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			if rec.Code != tc.expectedCode {
				t.Errorf("Expected status %d, got %d", tc.expectedCode, rec.Code)
			}

			if rec.Body.String() != tc.expectedBody {
				t.Errorf("Expected body %s, got %s", tc.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestWithRealServer(t *testing.T) {
	server := httptest.NewServer(setupServer())
	defer server.Close()

	resp, err := http.Get(server.URL + "/home")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %d", resp.StatusCode)
	}

	expectedBody := "Hello Home\n"
	body := make([]byte, len(expectedBody))
	resp.Body.Read(body)
	if string(body) != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, string(body))
	}
}
