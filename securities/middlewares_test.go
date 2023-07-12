package securities

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// TestCoreAuthentication tests the middleware of api key header auth.
func TestCoreAuthenticate(t *testing.T) {
	t.Setenv("API_KEY", "someGoodAPIKey")
	t.Run("ErrorWrongKey", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := &http.Request{
			URL:    &url.URL{},
			Header: make(http.Header),
		}

		c.Request = req
		CoreAuthenticate(c)
		assert.Equal(t, 401, w.Code)
	})

	t.Run("CorrectAPIKey", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := &http.Request{
			URL:    &url.URL{},
			Header: make(http.Header),
		}

		req.Header.Add("X-API-Key", os.Getenv("API_KEY"))
		c.Request = req
		CoreAuthenticate(c)
		assert.Equal(t, 200, w.Code)
	})
}
