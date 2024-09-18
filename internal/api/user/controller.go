package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
	"ocserv/pkg/password"
	"strconv"
	"time"
)

type Controller struct {
	userRepository  repository.UserRepositoryInterface
	tokenRepository repository.TokenRepositoryInterface
}

func NewUserController() *Controller {
	return &Controller{
		userRepository: repository.NewUserRepository(),
	}
}

// CreateAdminUser godoc
// @Summary      Set up an admin user
// @Description  Set up an admin or superuser during site initialization
// @Tags         user
// @Produce      json
// @Param        user  body      CreateData  true  "Request Body"
// @Success      200  {object}  CreateResponse
// @Failure      400  {object}  nil
// @Router       /api/v1/user/ [post]
func (controller *Controller) CreateAdminUser(c *gin.Context) {
	var data CreateData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := controller.userRepository.Exists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin user already exists"})
		return
	}

	user := models.User{
		Username: data.Username,
		Password: password.MakeHash(data.Password),
		IsAdmin:  true,
	}
	newUser, err := controller.userRepository.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Admin:    newUser.IsAdmin,
	})
}

// Login godoc
// @Summary      Login
// @Description  Login admin or staff user to get token
// @Tags         user
// @Produce      json
// @Param        user  body CreateLoginData  true  "Request Body"
// @Success      201 {object} LoginResponse
// @Failure      400  {object}  nil
// @Router       /api/v1/user/login/ [post]
func (controller *Controller) Login(c *gin.Context) {
	var data CreateLoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, ok := Authenticate(controller.userRepository, data)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	expireAt := time.Now().Add(24 * time.Hour).Unix()
	if data.RememberMe {
		expireAt = time.Now().Add(30 * 24 * time.Hour).Unix()
	}
	tokenObj := models.Token{
		UserID:   user.ID,
		ExpireAt: expireAt,
	}

	token, err := controller.tokenRepository.Create(&tokenObj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, LoginResponse{
		Token:    token.Key,
		ExpireAt: token.ExpireAt,
	})
}

// UpdatePassword godoc
// @Summary      Update Password
// @Description  Update admin or staff user password (self change)
// @Tags         user
// @Produce      json
// @Param        user  body UpdateData  true  "Request Body"
// @Param        Authorization header string true "Bearer token"
// @Success      202
// @Failure      400  {object}  nil
// @Router       /api/v1/user/password/ [post]
func (controller *Controller) UpdatePassword(c *gin.Context) {
	var data UpdateData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.CurrentPassword != data.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "old password and new password do not match"})
		return
	}

	userContext, _ := c.Get("user")
	user := userContext.(models.User)

	if checkPassword := password.Compare(data.CurrentPassword, user.Password); !checkPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid old password"})
		return
	}

	err = controller.userRepository.UpdatePassword(user.ID, data.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

// CreateStaff godoc
// @Summary      Create staff user
// @Description  Create staff user
// @Tags         user
// @Produce      json
// @Param        user  body CreateStaffData  true  "Request Body"
// @Param        Authorization header string true "Bearer token"
// @Success      201 {object} CreateResponse
// @Failure      403 {object} nil "Admin Permission required"
// @Failure      400  {object}  nil
// @Router       /api/v1/user/staffs/ [post]
func (controller *Controller) CreateStaff(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin Permission required"})
		return
	}

	var data CreateStaffData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username: data.Username,
		Password: password.MakeHash(data.Password),
		IsAdmin:  false,
	}

	newUser, err := controller.userRepository.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Admin:    newUser.IsAdmin,
	})
}

// UpdateStaffPassword godoc
// @Summary      Update staff password
// @Description  Update staff user password(by admin)
// @Tags         user
// @Produce      json
// @Param        user  body UpdateData  true  "Request Body"
// @Param        Authorization header string true "Bearer token"
// @Success      202
// @Failure      400  {object}  nil
// @Failure      403 {object} nil "Admin Permission required"
// @Router       /api/v1/user/staffs/:id/password/ [post]
func (controller *Controller) UpdateStaffPassword(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin Permission required"})
		return
	}

	var data UpdateData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.CurrentPassword != data.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "old password and new password do not match"})
		return
	}

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	user, err := controller.userRepository.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = controller.userRepository.UpdatePassword(user.ID, data.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

// DeleteStaff godoc
// @Summary      Delete staff
// @Description  Delete staff user(by admin)
// @Tags         user
// @Produce      json
// @Param        user  body UpdateData  true  "Request Body"
// @Param        Authorization header string true "Bearer token"
// @Success      204
// @Failure      404  {object}  nil
// @Failure      403 {object} nil "Admin Permission required"
// @Router       /api/v1/user/staffs/:id/ [delete]
func (controller *Controller) DeleteStaff(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin Permission required"})
		return
	}

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := controller.userRepository.DeleteStaffUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
