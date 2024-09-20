package site

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"ocserv/internal/api/site"
	"ocserv/internal/models"
	"ocserv/internal/providers/routes/middleware"
	"ocserv/pkg/config"
	"ocserv/pkg/database"
	"ocserv/pkg/routing"
	"ocserv/pkg/testutils"
	"os"
	"strings"
	"testing"
)

var (
	siteController *site.Controller
	router         *gin.Engine
	adminUser      *models.User
	adminToken     string
)

func init() {
	testutils.LoadTestEnv()
	config.Set()
	database.Connect()
	routing.Init()
	siteController = site.NewSiteController()
	router = routing.GetRouter()
	addRoutes()
	adminUser = testutils.CreateTestAdminUser()
	adminToken = testutils.CreateTestAdminToken(adminUser.ID)
}

func addRoutes() {
	router.GET("/api/v1/site/", siteController.Get)
	router.POST("/api/v1/site/", siteController.Create)
	router.PATCH("/api/v1/site/", middleware.TokenMiddleware(), siteController.Update)
}

func TestMain(m *testing.M) {
	code := m.Run()
	defer testutils.DeleteTestAdminUser()
	os.Exit(code)
}

func TestSiteCreate(t *testing.T) {
	var data map[string]interface{}

	body := map[string]interface{}{
		"captcha_site_key":   "",
		"captcha_secret_key": "",
		"default_traffic":    10,
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(&body)
	req, _ := http.NewRequest("POST", "/api/v1/site/", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	err := json.Unmarshal(w.Body.Bytes(), &data)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, data["id"])
	assert.Equal(t, "", data["captcha_site_key"])
	assert.Equal(t, "", data["captcha_secret_key"])
	assert.Equal(t, float64(10), data["default_traffic"])

	// recreate test to get exists error
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestSiteGet(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/site/", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.Nil(t, err)
	assert.Equal(t, float64(1), responseData["id"])
	assert.Equal(t, float64(10), responseData["default_traffic"])
}

func TestSiteUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	body := map[string]interface{}{
		"default_traffic": 20,
	}
	jsonBody, _ := json.Marshal(&body)
	req, _ := http.NewRequest("PATCH", "/api/v1/site/", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))

	router.ServeHTTP(w, req)
	assert.Equal(t, 202, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/site/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.Nil(t, err)
	assert.Equal(t, float64(1), responseData["id"])
	assert.Equal(t, float64(20), responseData["default_traffic"])
}
