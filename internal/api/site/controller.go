package site

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
)

type Controller struct {
	siteRepository repository.SiteRepositoryInterface
}

func NewSiteController() *Controller {
	return &Controller{
		siteRepository: repository.NewSiteRepository(),
	}
}

// Get godoc
// @Summary      Get site configuration
// @Description  Get site configuration
// @Tags         site
// @Produce      json
// @Success      200  {object}  models.Site
// @Failure      400  {object}  nil
// @Router       /api/v1/site/ [get]
func (controller *Controller) Get(c *gin.Context) {
	site, err := controller.siteRepository.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, site)
}

// Create godoc
// @Summary      Post site configuration
// @Description  Post site configuration
// @Tags         site
// @Produce      json
// @Param        site  body      Data  true  "Request Body"
// @Success      200  {object}  models.Site
// @Failure      400  {object}  nil
// @Router       /api/v1/site/ [post]
func (controller *Controller) Create(c *gin.Context) {
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.DefaultTraffic <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid default_traffic"})
		return
	}
	site := models.Site{
		CaptchaSiteKey:   data.CaptchaSiteKey,
		CaptchaSecretKey: data.CaptchaSecretKey,
		DefaultTraffic:   data.DefaultTraffic,
	}
	err := controller.siteRepository.Create(&site)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": site})
}

// Update godoc
// @Summary      Update site configuration
// @Description  Update site configuration
// @Tags         site
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        site  body     DataUpdate  false  "Request Body"
// @Success      200  {object}  models.Site
// @Failure      400  {object}  nil
// @Failure      401  {object}  nil
// @Router       /api/v1/site/ [patch]
func (controller *Controller) Update(c *gin.Context) {
	if isStaff, exists := c.Get("isStaff"); !exists || !isStaff.(bool) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "only admin can update site configs",
		})
		return
	}

	var data DataUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var site models.Site
	if data.DefaultTraffic <= 0 {
		site.DefaultTraffic = data.DefaultTraffic
	}
	if data.CaptchaSiteKey != "" {
		site.CaptchaSiteKey = data.CaptchaSiteKey
	}
	if data.CaptchaSecretKey != "" {
		site.CaptchaSecretKey = data.CaptchaSecretKey
	}
	update, err := controller.siteRepository.Update(&site)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, update)
}
