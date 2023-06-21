package global

import (
	"blog-server/pkg/logger"
	"blog-server/pkg/setting"
)

// 全局变量
// global variable will update with develope
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
