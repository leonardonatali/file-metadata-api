package test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/server"
	"github.com/leonardonatali/file-metadata-api/pkg/storage"
)

func Test_Integration(t *testing.T) {
	cfg := &config.Config{}
	storageCfg := &storage.StorageConfig{}

	if err := cfg.Load(); err != nil {
		log.Fatalf("cannot load app config: %s", err.Error())
	}

	if err := storageCfg.Load(); err != nil {
		log.Fatalf("cannot load storage config: %s", err.Error())
	}

	s := server.NewServer(cfg, storageCfg)
	s.Setup()

	tests := []struct {
		name                    string
		endpoint                string
		method                  string
		requestBody             io.Reader
		expectedStatusCode      int
		expectedResponseContent string
	}{
		{
			name:                    "GET /ok",
			endpoint:                "/ok",
			method:                  http.MethodGet,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "GET /files/filetree",
			endpoint:                "/files/filetree",
			method:                  http.MethodGet,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "POST /files/upload",
			endpoint:                "/files/upload",
			method:                  http.MethodPost,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "GET /files/:id/download",
			endpoint:                "/files/:id/download",
			method:                  http.MethodGet,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "GET /files/:id/metadata",
			endpoint:                "/files/:id/metadata",
			method:                  http.MethodGet,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "DELETE /files/:id",
			endpoint:                "/files/:id",
			method:                  http.MethodDelete,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "PATCH /files/:id",
			endpoint:                "/files/:id",
			method:                  http.MethodPatch,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "PUT /files/:id",
			endpoint:                "/files/:id",
			method:                  http.MethodPut,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
		{
			name:                    "GET /files/",
			endpoint:                "/files",
			method:                  http.MethodGet,
			requestBody:             nil,
			expectedStatusCode:      100,
			expectedResponseContent: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.endpoint, tt.requestBody)

			testHTTPResponse(t, s.Router, req, func(w *httptest.ResponseRecorder) {

				statusIsOk := w.Code == tt.expectedStatusCode
				content, err := io.ReadAll(w.Body)

				if err != nil {
					t.Errorf("IntegrationTests(%s)\n%s %s\nError: %s", tt.method, tt.endpoint, tt.name, err.Error())
				}

				if !statusIsOk {
					t.Errorf("IntegrationTests(%s)\n%s %s\nStatus is not ok: expected %d, received %d", tt.method, tt.endpoint, tt.name, tt.expectedStatusCode, w.Code)
				}

				if string(content) != tt.expectedResponseContent {
					t.Errorf("IntegrationTests(%s)\n%s %s\nresponseis not ok: \nexpected %s\nreceived %s", tt.method, tt.endpoint, tt.name, tt.expectedResponseContent, string(content))
				}
			})

		})
	}
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder)) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	f(w)
}
