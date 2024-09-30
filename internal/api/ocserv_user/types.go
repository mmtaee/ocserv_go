package ocserv

import "ocserv/internal/models"

type CreateOcservUserBody struct {
	Group          string                 `json:"group" binding:"required"`
	Username       string                 `json:"username" binding:"required,min=4,max=8"`
	Password       string                 `json:"password" binding:"omitempty,min=4,max=12"`
	IsActive       bool                   `json:"is_active" binding:"omitempty,oneof=true false"`
	ExpireAt       int64                  `json:"expire_at" binding:"omitempty"`
	DefaultTraffic float64                `json:"default_traffic" binding:"required,gt=0"`
	TrafficType    models.ServiceTypeEnum `json:"traffic_type" binding:"required,oneof=FREE MONTHLY TOTALLY"`
}

type UpdateOcservUserBody struct {
	Group          string                 `json:"group"`
	Password       string                 `json:"password" binding:"omitempty,min=4,max=12"`
	IsActive       bool                   `json:"is_active" binding:"omitempty,oneof=true false"`
	ExpireAt       int64                  `json:"expire_at" binding:"omitempty"`
	DefaultTraffic float64                `json:"default_traffic" binding:"omitempty,gt=0"`
	TrafficType    models.ServiceTypeEnum `json:"traffic_type" binding:"omitempty,oneof=FREE MONTHLY TOTALLY"`
}
