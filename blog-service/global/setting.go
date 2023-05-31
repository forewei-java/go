package global

import (
	"github.com/forewei-java/go/blog-service/pkg/logger"
	"github.com/forewei-java/go/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
