package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
	"ocserv/pkg/password"
)

type Controller struct {
	userRepository repository.UserRepositoryInterface
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
// @Success      200  {object}  models.User
// @Failure      400  {object}  nil
// @Router       /api/v1/user/ [post]
func (controller *Controller) CreateAdminUser(c *gin.Context) {
	var data CreateData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	c.JSON(http.StatusCreated, newUser)
}

// UpdatePassword godoc
// @Summary      Update Password
// @Description  Update admin or staff user password
// @Tags         user
// @Produce      json
// @Param        user  body UpdateData  true  "Request Body"
// @Success      202
// @Failure      400  {object}  nil
// @Router       /api/v1/user/password/ [patch]
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
