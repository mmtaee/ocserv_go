package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"ocserv/internal/api/user"
	"ocserv/internal/models"
	"ocserv/internal/providers/routes/middleware"
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
	adminPassword  string
)

func init() {
	testutils.LoadTestEnv()
	config.Set()
	database.Connect()
	testutils.DeleteTestAdminUser()
	routing.Init()
	userController = user.NewUserController()
	router = routing.GetRouter()
	addRoutes()
	adminPassword = "test-admin-password"
	adminUser = testutils.CreateTestAdminUser()
	adminToken = testutils.CreateTestAdminToken(adminUser.ID)
}

func addRoutes() {
	router.POST("/api/v1/users/", userController.CreateAdminUser)
	router.POST("/api/v1/users/login/", userController.Login)
	router.PATCH("/api/v1/users/password/", middleware.TokenMiddleware(), userController.UpdatePassword)

	staffGroup := router.Group("/api/v1/users/staffs/", middleware.TokenMiddleware())
	staffGroup.POST("/", userController.CreateStaff)
	staffGroup.PATCH("/:id/password/", userController.UpdateStaffPassword)
	staffGroup.DELETE("/:id/", userController.DeleteStaff)
}

func TestMain(m *testing.M) {
	testutils.DeleteTestAdminUser()
	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	var data map[string]interface{}

	body := map[string]interface{}{
		"username": "test-admin-user-api",
		"password": adminPassword,
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
	assert.NotEqual(t, data["password"], adminPassword)
}

func TestLogin(t *testing.T) {
	var data map[string]interface{}
	body := map[string]interface{}{
		"username":    "test-admin-user-api",
		"password":    adminPassword,
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
	assert.LessOrEqual(t, int64(data["expire_at"].(float64)), time.Now().Add(24*30*time.Hour).Unix())
	adminToken = data["token"].(string)
}

func TestUpdatePassword(t *testing.T) {
	body := map[string]interface{}{
		"current_password": adminPassword,
		"new_password":     "test-admin-password-update",
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(&body)
	req, _ := http.NewRequest("PATCH", "/api/v1/users/password/", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))

	router.ServeHTTP(w, req)
	t.Log(w.Body.String())
	assert.Equal(t, 202, w.Code)
}

func TestCreateStaff(t *testing.T) {
	var data map[string]interface{}
	body := map[string]interface{}{
		"username": "test-staff-user-api",
		"password": "test-staff-user-password",
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(&body)
	req, _ := http.NewRequest("POST", "/api/v1/users/staffs/", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	err := json.Unmarshal(w.Body.Bytes(), &data)
	assert.Nil(t, err)
	assert.NotEqual(t, data["id"], 0)
	assert.Equal(t, data["username"], "test-staff-user-api")
	assert.False(t, data["is_admin"].(bool))
}

func TestDeleteStaff(t *testing.T) {
	staffUser := testutils.CreateTestStaffUser()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/users/staffs/%d/", staffUser.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
}

func TestUpdateStaffPassword(t *testing.T) {
	staffUser := testutils.CreateTestStaffUser()
	body := map[string]interface{}{
		"password": "test-staff-user-password-update",
	}

	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(&body)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/v1/users/staffs/%d/password/", staffUser.ID),
		strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))

	router.ServeHTTP(w, req)
	assert.Equal(t, 202, w.Code)
}
