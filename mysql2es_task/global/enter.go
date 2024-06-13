package global

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mysql2es_task/config"
)

var (
	Config   *config.Config
	ESClient *elastic.Client
	Logrus   *logrus.Logger
	DB       *gorm.DB
)
