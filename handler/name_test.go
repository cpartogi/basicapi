package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetDataNotAllowed(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		want       string
		statusCode int
	}{
		{
			name:       "with bad method",
			method:     http.MethodPost,
			want:       "{\"code\":405,\"message\":\"method not allowed\"}",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			request := httptest.NewRequest(tc.method, "/", nil)
			responseRecorder := httptest.NewRecorder()

			GetData(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

func TestGetDataNotExist(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		want       string
		statusCode int
	}{
		{
			name:       "resource not exist",
			method:     http.MethodGet,
			want:       "{\"code\":404,\"message\":\"resource with ID 1 not exist\"}",
			statusCode: http.StatusNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			request := httptest.NewRequest(tc.method, "/?id=1", nil)
			responseRecorder := httptest.NewRecorder()

			GetData(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

func TestGetDataAll(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		want       string
		statusCode int
	}{
		{
			name:       "show all data",
			method:     http.MethodGet,
			want:       "{\"code\":200,\"data\":null}",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			request := httptest.NewRequest(tc.method, "/", nil)
			responseRecorder := httptest.NewRecorder()

			GetData(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

func TestGetDataInvalidId(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		want       string
		statusCode int
	}{
		{
			name:       "show invalid id",
			method:     http.MethodGet,
			want:       "{\"code\":400,\"message\":\"invalid or empty ID: a\"}",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			request := httptest.NewRequest(tc.method, "/?id=a", nil)
			responseRecorder := httptest.NewRecorder()

			GetData(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

func TestGetDataWithIds(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		want       string
		statusCode int
	}{
		{
			name:       "show data with ids",
			method:     http.MethodGet,
			want:       "{\"code\":200,\"data\":null}",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			request := httptest.NewRequest(tc.method, "/?id=1,2,3", nil)
			responseRecorder := httptest.NewRecorder()

			GetData(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
