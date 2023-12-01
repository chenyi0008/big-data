package svc

import (
	"bigdata/core/internal/config"
	"bigdata/core/models"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		DB:     models.InitMysql(c.Mysql.DataSource),
	}
}
