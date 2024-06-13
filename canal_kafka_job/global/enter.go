package global

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var (
	ESClient *elastic.Client
	Logger   *logrus.Logger
)
