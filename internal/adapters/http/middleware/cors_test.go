package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name            string
		allowedOrigins  string
		requestOrigin   string
		expectedOrigin  string
		expectedVary    string
		shouldSetOrigin bool
	}{
		{
			name:            "wildcard allows any origin",
			allowedOrigins:  "*",
			requestOrigin:   "https://example.com",
			expectedOrigin:  "*",
			expectedVary:    "",
			shouldSetOrigin: true,
		},
		{
			name:            "wildcard with no origin header",
			allowedOrigins:  "*",
			requestOrigin:   "",
			expectedOrigin:  "*",
			expectedVary:    "",
			shouldSetOrigin: true,
		},
		{
			name:            "single allowed origin matches",
			allowedOrigins:  "https://example.com",
			requestOrigin:   "https://example.com",
			expectedOrigin:  "https://example.com",
			expectedVary:    "Origin",
			shouldSetOrigin: true,
		},
		{
			name:            "single allowed origin does not match",
			allowedOrigins:  "https://example.com",
			requestOrigin:   "https://evil.com",
			expectedOrigin:  "",
			expectedVary:    "",
			shouldSetOrigin: false,
		},
		{
			name:            "multiple allowed origins - first matches",
			allowedOrigins:  "https://example.com,https://app.example.com",
			requestOrigin:   "https://example.com",
			expectedOrigin:  "https://example.com",
			expectedVary:    "Origin",
			shouldSetOrigin: true,
		},
		{
			name:            "multiple allowed origins - second matches",
			allowedOrigins:  "https://example.com,https://app.example.com",
			requestOrigin:   "https://app.example.com",
			expectedOrigin:  "https://app.example.com",
			expectedVary:    "Origin",
			shouldSetOrigin: true,
		},
		{
			name:            "multiple allowed origins - none match",
			allowedOrigins:  "https://example.com,https://app.example.com",
			requestOrigin:   "https://evil.com",
			expectedOrigin:  "",
			expectedVary:    "",
			shouldSetOrigin: false,
		},
		{
			name:            "multiple allowed origins with spaces",
			allowedOrigins:  "https://example.com, https://app.example.com",
			requestOrigin:   "https://app.example.com",
			expectedOrigin:  "https://app.example.com",
			expectedVary:    "Origin",
			shouldSetOrigin: true,
		},
		{
			name:            "no origin header with specific origins",
			allowedOrigins:  "https://example.com",
			requestOrigin:   "",
			expectedOrigin:  "",
			expectedVary:    "",
			shouldSetOrigin: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("GET", "/test", nil)
			if tt.requestOrigin != "" {
				c.Request.Header.Set("Origin", tt.requestOrigin)
			}

			// Create middleware
			handler := CORS(tt.allowedOrigins)

			// Execute
			handler(c)

			// Assert
			if tt.shouldSetOrigin {
				assert.Equal(t, tt.expectedOrigin, w.Header().Get("Access-Control-Allow-Origin"))
				if tt.expectedVary != "" {
					assert.Equal(t, tt.expectedVary, w.Header().Get("Vary"))
				}
			} else {
				assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
			}

			// Verify common CORS headers are always set
			assert.Equal(t, "GET", w.Header().Get("Access-Control-Allow-Methods"))
			assert.Equal(t, "Content-Type, Authorization", w.Header().Get("Access-Control-Allow-Headers"))
		})
	}
}

func TestCORS_PreflightRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		allowedOrigins string
		requestOrigin  string
		expectedStatus int
	}{
		{
			name:           "OPTIONS request with wildcard",
			allowedOrigins: "*",
			requestOrigin:  "https://example.com",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "OPTIONS request with allowed origin",
			allowedOrigins: "https://example.com",
			requestOrigin:  "https://example.com",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "OPTIONS request with disallowed origin",
			allowedOrigins: "https://example.com",
			requestOrigin:  "https://evil.com",
			expectedStatus: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("OPTIONS", "/test", nil)
			c.Request.Header.Set("Origin", tt.requestOrigin)

			// Create middleware
			handler := CORS(tt.allowedOrigins)

			// Execute
			handler(c)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.True(t, c.IsAborted())
		})
	}
}

func TestCORS_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name            string
		allowedOrigins  string
		requestOrigin   string
		expectedOrigin  string
		shouldSetOrigin bool
	}{
		{
			name:            "production with single origin",
			allowedOrigins:  "https://myapp.com",
			requestOrigin:   "https://myapp.com",
			expectedOrigin:  "https://myapp.com",
			shouldSetOrigin: true,
		},
		{
			name:            "production with multiple origins",
			allowedOrigins:  "https://myapp.com,https://staging.myapp.com",
			requestOrigin:   "https://staging.myapp.com",
			expectedOrigin:  "https://staging.myapp.com",
			shouldSetOrigin: true,
		},
		{
			name:            "development with wildcard",
			allowedOrigins:  "*",
			requestOrigin:   "http://localhost:3000",
			expectedOrigin:  "*",
			shouldSetOrigin: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup router with middleware
			router := gin.New()
			router.Use(CORS(tt.allowedOrigins))
			router.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			})

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/test", nil)
			if tt.requestOrigin != "" {
				req.Header.Set("Origin", tt.requestOrigin)
			}

			// Execute
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, http.StatusOK, w.Code)
			if tt.shouldSetOrigin {
				assert.Equal(t, tt.expectedOrigin, w.Header().Get("Access-Control-Allow-Origin"))
			} else {
				assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
			}
		})
	}
}
