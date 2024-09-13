package repository

import (
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/database"
)

type SiteRepository struct {
	db *gorm.DB
}

type SiteRepositoryInterface interface {
	Get() (*models.Site, error)
	Create(*models.Site) error
	Update(*models.Site) (*models.Site, error)
}

func NewSiteRepository() *SiteRepository {
	return &SiteRepository{
		db: database.Connection(),
	}
}

func (s *SiteRepository) Get() (*models.Site, error) {
	ch := make(chan struct {
		site *models.Site
		err  error
	}, 1)

	go func() {
		var site *models.Site
		err := s.db.First(&site).Error
		ch <- struct {
			site *models.Site
			err  error
		}{site, err}
	}()
	result := <-ch
	return result.site, result.err
}

func (s *SiteRepository) Create(site *models.Site) (err error) {
	return s.db.Create(site).Error
}

func (s *SiteRepository) Update(site *models.Site) (*models.Site, error) {
	ch := make(chan struct {
		site *models.Site
		err  error
	})
	go func() {
		err := s.db.Save(site).Error
		ch <- struct {
			site *models.Site
			err  error
		}{site, err}
	}()

	result := <-ch
	return result.site, result.err
}
