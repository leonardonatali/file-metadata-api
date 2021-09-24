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
			expectedResponseContent: `{"error":"the request input must be multipart/form"}`,
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
				"File": GetFile(),
			},
			expectedStatusCode:      http.StatusCreated,
			expectedResponseContent: ``,
		},
		{
			name:                    "Should return a correct file tree",
			endpoint:                "/files/filetree",
			method:                  http.MethodGet,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusOK,
			expectedResponseContent: `[{"CurrentDir":"","Children":[{"CurrentDir":"path","Children":[{"CurrentDir":"test"}]}]}]`,
		},
		{
			name:                    "Should return 404 for non existing file",
			endpoint:                "/files/999/download",
			method:                  http.MethodGet,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusNotFound,
			expectedResponseContent: `{"error":"file not found"}`,
		},
		{
			name:               "Should return correct download link",
			endpoint:           "/files/1/download",
			method:             http.MethodGet,
			headers:            getAuthHeader(),
			form:               map[string]io.Reader{},
			expectedStatusCode: http.StatusOK,
			//expectedResponseContent: `{"DownloadURL":"http://minio:9000/app-files-test/1/path/test/1_file.txt?"}`,
		},
		{
			name:                    "Should return 404 for metadata of not found file",
			endpoint:                "/files/999/metadata",
			method:                  http.MethodGet,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusNotFound,
			expectedResponseContent: `{"error":"file not found"}`,
		},
		{
			name:                    "Should return file metadata",
			endpoint:                "/files/1/metadata",
			method:                  http.MethodGet,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusOK,
			expectedResponseContent: `[{"ID":1,"File":null,"FileID":1,"Key":"filename","Value":"file.txt"},{"ID":2,"File":null,"FileID":1,"Key":"path","Value":"/path/test"},{"ID":3,"File":null,"FileID":1,"Key":"size","Value":"16"},{"ID":4,"File":null,"FileID":1,"Key":"type","Value":"application/octet-stream"}]`,
		},
		{
			name:                    "Should return validation error for path",
			endpoint:                "/files/999",
			method:                  http.MethodPatch,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusBadRequest,
			expectedResponseContent: `{"error":"Key: 'UpdateFilePathDto.Path' Error:Field validation for 'Path' failed on the 'required' tag"}`,
		},
		{
			name:     "Should return 404 for not found product to update",
			endpoint: "/files/999",
			method:   http.MethodPatch,
			headers:  getAuthHeader(),
			form: map[string]io.Reader{
				"Path": strings.NewReader("/path/updated"),
			},
			expectedStatusCode:      http.StatusNotFound,
			expectedResponseContent: `{"error":"file not found"}`,
		},
		{
			name:     "Should update product path",
			endpoint: "/files/1",
			method:   http.MethodPatch,
			headers:  getAuthHeader(),
			form: map[string]io.Reader{
				"Path": strings.NewReader("/path/updated"),
			},
			expectedStatusCode:      http.StatusOK,
			expectedResponseContent: ``,
		},
		{
			name:                    "Should fails at validation request type in full update file",
			endpoint:                "/files/1",
			method:                  http.MethodPut,
			headers:                 getAuthHeader(),
			form:                    nil,
			expectedStatusCode:      http.StatusUnprocessableEntity,
			expectedResponseContent: `{"error":"the request input must be multipart/form"}`,
		},
		{
			name:                    "Should fails at validation of Path in full update file",
			endpoint:                "/files/1",
			method:                  http.MethodPut,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{"type": strings.NewReader("test")},
			expectedStatusCode:      http.StatusBadRequest,
			expectedResponseContent: `{"error":"Key: 'UpdateFileDto.Path' Error:Field validation for 'Path' failed on the 'required' tag\nKey: 'UpdateFileDto.File' Error:Field validation for 'File' failed on the 'required' tag"}`,
		},
		{
			name:                    "Should fails at validation file form param in full update file",
			endpoint:                "/files/1",
			method:                  http.MethodPut,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{"Path": strings.NewReader("/path/test/full/updated")},
			expectedStatusCode:      http.StatusBadRequest,
			expectedResponseContent: `{"error":"Key: 'UpdateFileDto.File' Error:Field validation for 'File' failed on the 'required' tag"}`,
		},
		{
			name:     "Should return file not found error at full update",
			endpoint: "/files/999",
			method:   http.MethodPut,
			headers:  getAuthHeader(),
			form: map[string]io.Reader{
				"Path": strings.NewReader("/path/test"),
				"File": GetFile(),
			},
			expectedStatusCode:      http.StatusNotFound,
			expectedResponseContent: `{"error":"file not found"}`,
		},
		{
			name:     "Should full update a file",
			endpoint: "/files/1",
			method:   http.MethodPut,
			headers:  getAuthHeader(),
			form: map[string]io.Reader{
				"Path": strings.NewReader("/path/test"),
				"File": GetFile(),
			},
			expectedStatusCode:      http.StatusOK,
			expectedResponseContent: ``,
		},
		{
			name:                    "Should return all files",
			endpoint:                "/files",
			method:                  http.MethodGet,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusOK,
			expectedResponseContent: "",
		},
		{
			name:                    "Should return 404 for not found product to delete",
			endpoint:                "/files/999",
			method:                  http.MethodDelete,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusNotFound,
			expectedResponseContent: `{"error":"file not found"}`,
		},
		{
			name:                    "Should delete product",
			endpoint:                "/files/1",
			method:                  http.MethodDelete,
			headers:                 getAuthHeader(),
			form:                    map[string]io.Reader{},
			expectedStatusCode:      http.StatusOK,
			expectedResponseContent: ``,
		},
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

				if tt.expectedResponseContent != "" && strings.TrimSpace(string(content)) != strings.TrimSpace(tt.expectedResponseContent) {
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
