package ocserv

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
	customErrors "ocserv/pkg/errors"
	"strconv"
)

type Controller struct {
	ocservRepository repository.OcservRepositoryInterface
}

func NewOcservController() *Controller {
	return &Controller{
		ocservRepository: repository.NewOcservRepository(),
	}
}

// Create godoc
// @Summary      Create Ocserv User
// @Description  Create Ocserv User
// @Tags          ocserv user
// @Produce      json
// @Param        site  body     CreateOcservUserBody  true  "Request Body"
// @Success      201  {object}  models.OcservUser
// @Failure      400  {object}  nil
// @Router       /api/v1/ocserv/ [post]
func (controller *Controller) Create(c *gin.Context) {
	var data CreateOcservUserBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.InvalidBodyError(err))
		return
	}
	user := models.OcservUser{
		Username:       data.Username,
		Password:       data.Password,
		IsActive:       data.IsActive,
		DefaultTraffic: data.DefaultTraffic,
		TrafficType:    data.TrafficType,
		ExpireAt:       data.ExpireAt,
	}
	createdUser, err := controller.ocservRepository.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

// Update godoc
// @Summary      Update Ocserv User
// @Description  Update Ocserv User
// @Tags         ocserv user
// @Produce      json
// @Param        site  body     UpdateOcservUserBody  true  "Request Body"
// @Success      200  {object}  models.OcservUser
// @Failure      400  {object}  nil
// @Failure      404  {object}  nil
// @Router       /api/v1/ocserv/:id/ [patch]
func (controller *Controller) Update(c *gin.Context) {
	var data UpdateOcservUserBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.InvalidBodyError(err))
		return
	}
	ocservUserIDStr := c.Param("id")
	ocservUserID, err := strconv.Atoi(ocservUserIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ocserv user id"})
		return
	}
	ocservUser, err := controller.ocservRepository.GetUserByID(ocservUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ocserv user not found"})
		return
	}
	if data.Group != "" {
		ocservUser.Group = data.Group
	}
	if data.Password != "" {
		ocservUser.Password = data.Password
	}
	ocservUser.IsActive = data.IsActive
	if data.ExpireAt != 0 {
		ocservUser.ExpireAt = data.ExpireAt
	}
	if data.DefaultTraffic > 0 {
		ocservUser.DefaultTraffic = data.DefaultTraffic
	}
	if data.TrafficType != "" {
		ocservUser.TrafficType = data.TrafficType
	}
	createdUser, err := controller.ocservRepository.UpdateUser(ocservUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

// Delete godoc
// @Summary      Delete ocserv user
// @Description  Delete ocserv user
// @Tags         ocserv user
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Success      204
// @Failure      400  {object}  nil
// @Failure      404  {object}  nil
// @Router       /api/v1/ocserv/:id/ [delete]
func (controller *Controller) Delete(c *gin.Context) {
	ocservUserIDStr := c.Param("id")
	ocservUserID, err := strconv.Atoi(ocservUserIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ocserv user id"})
		return
	}
	ocservUser, err := controller.ocservRepository.GetUserByID(ocservUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ocserv user not found"})
		return
	}
	err = controller.ocservRepository.DeleteUser(ocservUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Disconnect godoc
// @Summary      Disconnect ocserv user
// @Description  Disconnect ocserv user
// @Tags         ocserv user
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Success      202  {object}  nil
// @Failure      400  {object}  nil
// @Failure		 404  {object}  nil
// @Router       /api/v1/ocserv/:id/disconnect/ [post]
func (controller *Controller) Disconnect(c *gin.Context) {
	ocservUserIDStr := c.Param("id")
	ocservUserID, err := strconv.Atoi(ocservUserIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ocserv user id"})
		return
	}
	ocservUser, err := controller.ocservRepository.GetUserByID(ocservUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ocserv user not found"})
		return
	}
	log.Println(ocservUser)
	//	TODO: call ocserv

	c.JSON(http.StatusNoContent, nil)
}

func (controller *Controller) List(c *gin.Context) {
	//	TODO: add online status of users
}
