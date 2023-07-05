package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Setting is a global variable that stores the application configuration
type Setting struct {
	vp *viper.Viper
}

// NewSetting creates a new Setting instance.

// NewSetting initializes a new Setting instance.
//
// It accepts a variadic number of strings representing the paths to the configuration files.
// It returns a pointer to a Setting instance and an error if there was a problem reading the configuration file.
// NewSetting 初始化一个新的 Setting 实例。
//
// 它接受一个可变数量的字符串，表示配置文件的路径。
// 它返回一个指向 Setting 实例的指针，以及如果读取配置文件时出现问题则返回一个错误。

func NewSetting(configs ...string) (*Setting, error) {

	vp := viper.New()
	vp.SetConfigName("config")

	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

// WatchSettingChange 	use fsnotify update config
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection() //监听配置文件变化
		})
	}()
}
