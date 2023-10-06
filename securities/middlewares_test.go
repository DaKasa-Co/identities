package securities

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/DaKasa-Co/identities/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

// TestCoreAuthentication tests the middleware of api key header auth.
func TestCoreAuthenticate(t *testing.T) {
	t.Setenv("JWT_KEY", "someGoodJWTKey")
	t.Run("ErrorWrongJWTKey", func(t *testing.T) {
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

	t.Run("AuthSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := &http.Request{
			URL:    &url.URL{},
			Header: make(http.Header),
		}

		jwt, err := helper.GenerateJWT(uuid.New(), "usern", "", "", time.Now().AddDate(-16, 0, 0))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("X-JWT", jwt)
		c.Request = req
		CoreAuthenticate(c)
		assert.Equal(t, 200, w.Code)
	})
}
