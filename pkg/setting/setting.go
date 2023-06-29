package setting

import "github.com/spf13/viper"

// Setting is a global variable that stores the application configuration
type Setting struct {
	vp *viper.Viper
}

// NewSetting creates a new Setting instance.
//
// It takes a configPath string as a parameter and returns a pointer to a Setting struct
// and an error. The configPath parameter specifies the path to the configuration file.
// If the configPath is an empty string, the default "configs/" path is used.
// The function initializes a new viper instance, sets the configuration name to "config",
// adds the configPath to the configuration search paths, and sets the config type to "yaml".
// It then reads in the configuration file using vp.ReadInConfig() and returns the
// initialized Setting struct and any error encountered during the process.
func NewSetting(configPath string) (*Setting, error) {
	config := "configs/"
	if configPath != "" {
		config = configPath
	}
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(config)
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
