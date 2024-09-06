package user

import (
	"gorm.io/gorm"
	"ocserv/pkg/testutils"
	"testing"
)

var dbOcservUser *gorm.DB

func init() {
	dbOcservUser = testutils.GetTestDB()
}

func TestCreateOcservUser(t *testing.T) {}

func TestUpdateOcservUser(t *testing.T) {}

func TestDeleteOcservUser(t *testing.T) {}
