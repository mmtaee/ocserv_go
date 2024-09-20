package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"ocserv/internal/api/user"
	"ocserv/internal/models"
	"ocserv/pkg/config"
	"ocserv/pkg/database"
	"ocserv/pkg/routing"
	"ocserv/pkg/testutils"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	userController *user.Controller
	router         *gin.Engine
	adminUser      *models.User
	adminToken     string
)

func init() {
	testutils.LoadTestEnv()
	config.Set()
	database.Connect()
	routing.Init()
	userController = user.NewUserController()
	router = routing.GetRouter()
	addRoutes()
	adminUser = testutils.CreateTestAdminUser()
	adminToken = testutils.CreateTestAdminToken(adminUser.ID)
}

func addRoutes() {
	router.POST("/api/v1/users/", userController.CreateAdminUser)
	router.POST("/api/v1/users/login/", userController.Login)

}

func TestMain(m *testing.M) {
	testutils.DeleteTestAdminUser()
	code := m.Run()
	defer testutils.DeleteTestAdminUser()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	var data map[string]interface{}

	body := map[string]interface{}{
		"username": "test-admin-user-api",
		"password": "test-admin-password",
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(&body)
	req, _ := http.NewRequest("POST", "/api/v1/users/", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	err := json.Unmarshal(w.Body.Bytes(), &data)
	assert.Nil(t, err)

	assert.Equal(t, data["username"], "test-admin-user-api")
	assert.NotEqual(t, data["password"], "test-admin-password")
}

func TestLogin(t *testing.T) {
	var data map[string]interface{}
	body := map[string]interface{}{
		"username":    "test-admin-user-api",
		"password":    "test-admin-password",
		"remember_me": true,
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(&body)
	req, _ := http.NewRequest("POST", "/api/v1/users/login/", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	err := json.Unmarshal(w.Body.Bytes(), &data)
	assert.Nil(t, err)
	assert.NotEmpty(t, data["token"])
	assert.GreaterOrEqual(t, int64(data["expire_at"].(float64)), time.Now().Add(24*30*time.Hour).Unix())
}
