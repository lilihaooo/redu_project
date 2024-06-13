package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"redu/config"
)

var (
	Config      *config.Config
	DB          *gorm.DB
	Logrus      *logrus.Logger
	RedisClient *redis.Client
	ResMap      *config.ResMap
	ESClient    *elastic.Client
)
