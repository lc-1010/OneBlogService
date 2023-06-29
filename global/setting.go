package global

import (
	"github.com/lc-1010/OneBlogService/pkg/logger"
	"github.com/lc-1010/OneBlogService/pkg/setting"
)

// 全局变量
// global variable will update with develope
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSettings   *setting.EmailSettings
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
)
