package test

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/server"
	"github.com/leonardonatali/file-metadata-api/pkg/storage"
)

func Test_Integration(t *testing.T) {
	cfg := &config.Config{}
	cfg.Debug = false

	storageCfg := &storage.StorageConfig{}

	if err := cfg.Load(); err != nil {
		log.Fatalf("cannot load app config: %s", err.Error())
	}

	if err := storageCfg.Load(); err != nil {
		log.Fatalf("cannot load storage config: %s", err.Error())
	}

	s := server.NewServer(cfg, storageCfg)
	s.Db.Exec("DROP SCHEMA public CASCADE;")
	s.Db.Exec("CREATE SCHEMA public;")
	s.Setup()

	//Cria o arquivo para upload
	CreateFile()

	tests := []struct {
		name                    string
		endpoint                string
		headers                 http.Header
		method                  string
		form                    map[string]io.Reader
		expectedStatusCode      int
		expectedResponseContent string
	}{
		{
			name:                    "Should return 401 in GET /ok without token",
			endpoint:                "/ok",
			method:                  http.MethodGet,
			form:                    nil,
			expectedStatusCode:      http.StatusUnauthorized,
			expectedResponseContent: `{"error":"token must be provided"}`,
		},
		{
			name:                    "Should return 200 in GET /ok with token",
			endpoint:                "/ok",
			method:                  http.MethodGet,
			form:                    nil,
			headers:                 getAuthHeader(),
			expectedStatusCode:      http.StatusOK,
			expectedResponseContent: `:)`,
		},
		{
			name:                    "Should fails at validation request type in upload file",
			endpoint:                "/files/upload",
			method:                  http.MethodPost,
			headers:                 getAuthHeader(),
			form:                    nil,
			expectedStatusCode:      http.StatusUnprocessableEntity,
			expectedResponseContent: `{"error":"the request input must be multipStatusBadRequestart/form"}`,
		},
		{
			name:                    "Should fails at validation path form param in upload file",
			endpoint:                "/files/upload",
			method:                  http.MethodPost,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{"type": strings.NewReader("test")},
			expectedStatusCode:      http.StatusBadRequest,
			expectedResponseContent: `{"error":"Key: 'CreateFileDto.Path' Error:Field validation for 'Path' failed on the 'required' tag\nKey: 'CreateFileDto.File' Error:Field validation for 'File' failed on the 'required' tag"}`,
		},
		{
			name:                    "Should fails at validation file form param in upload file",
			endpoint:                "/files/upload",
			method:                  http.MethodPost,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{"Path": strings.NewReader("/path/test")},
			expectedStatusCode:      http.StatusBadRequest,
			expectedResponseContent: `{"error":"Key: 'CreateFileDto.File' Error:Field validation for 'File' failed on the 'required' tag"}`,
		},
		{
			name:     "Should upload file",
			endpoint: "/files/upload",
			method:   http.MethodPost,
			headers:  getAuthHeader(),
			form: map[string]io.Reader{
				"Path": strings.NewReader("/path/test"),
				"File": MustOpen(),
			},
			expectedStatusCode:      http.StatusCreated,
			expectedResponseContent: ``,
		},
		//{
		//	name:                    "GET /files/filetree",
		//	endpoint:                "/files/filetree",
		//	method:                  http.MethodGet,
		//	form:             map[string]io.Reader{},
		//	expectedStatusCode:      100,
		//	expectedResponseContent: "",
		//},
		//{
		//	name:                    "GET /files/:id/download",
		//	endpoint:                "/files/:id/download",
		//	method:                  http.MethodGet,
		//	form:             map[string]io.Reader{},
		//	expectedStatusCode:      100,
		//	expectedResponseContent: "",
		//},
		//{
		//	name:                    "GET /files/:id/metadata",
		//	endpoint:                "/files/:id/metadata",
		//	method:                  http.MethodGet,
		//	form:             map[string]io.Reader{},
		//	expectedStatusCode:      100,
		//	expectedResponseContent: "",
		//},
		//{
		//	name:                    "DELETE /files/:id",
		//	endpoint:                "/files/:id",
		//	method:                  http.MethodDelete,
		//	form:             map[string]io.Reader{},
		//	expectedStatusCode:      100,
		//	expectedResponseContent: "",
		//},
		//{
		//	name:                    "PATCH /files/:id",
		//	endpoint:                "/files/:id",
		//	method:                  http.MethodPatch,
		//	form:             map[string]io.Reader{},
		//	expectedStatusCode:      100,
		//	expectedResponseContent: "",
		//},
		//{
		//	name:                    "PUT /files/:id",
		//	endpoint:                "/files/:id",
		//	method:                  http.MethodPut,
		//	form:             map[string]io.Reader{},
		//	expectedStatusCode:      100,
		//	expectedResponseContent: "",
		//},
		//{
		//	name:                    "GET /files/",
		//	endpoint:                "/files",
		//	method:                  http.MethodGet,
		//	form:             map[string]io.Reader{},
		//	expectedStatusCode:      100,
		//	expectedResponseContent: "",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			content := &bytes.Buffer{}
			contentType := ""

			if len(tt.form) > 0 {
				content, contentType, _ = FormData(tt.form)
			}

			req, _ := http.NewRequest(tt.method, tt.endpoint, content)

			if len(tt.headers) > 0 {
				req.Header = tt.headers.Clone()
			}

			if contentType != "" {
				req.Header.Set("Content-Type", contentType)
			}

			testHTTPResponse(t, s.Router, req, func(w *httptest.ResponseRecorder) {

				statusIsOk := w.Code == tt.expectedStatusCode
				content, err := io.ReadAll(w.Body)

				if err != nil {
					t.Errorf("IntegrationTests\n%s\nError: %s", tt.name, err.Error())
				}

				if !statusIsOk {
					t.Errorf("IntegrationTests\n%s\nStatus is not ok: expected %d, received %d", tt.name, tt.expectedStatusCode, w.Code)
				}

				if string(content) != tt.expectedResponseContent {
					t.Errorf("IntegrationTests\n%s\nResponse is not ok: \nexpected %s\nreceived %s", tt.name, tt.expectedResponseContent, string(content))
				}
			})

		})
	}
}

func getAuthHeader() http.Header {
	header := http.Header{}
	header.Set("token", "123")
	return header
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder)) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	f(w)
}
