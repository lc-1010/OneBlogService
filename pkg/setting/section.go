package setting

import (
	"time"
)

// ServerSettingS Context from Code Snippet global/setting.go:package global
type ServerSettingS struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSettingS Context from Code Snippet global/setting.go:package global
type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	LogServePath          string
	LogFileName           string
	LogFileExt            string
	UploadSavePath        string
	UploadServerURL       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

// EmailSettings Context from Code Snippet global/setting.go:package global
type EmailSettings struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSl    bool
	From     string
	To       []string
}

// JWTSettingS Context from Code Snippet global/setting.go:package global
type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

// DatabaseSettingS Context from Code Snippet global/setting.go:package global
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
	// upload image
	UploadSavePath       string
	UploadServerURL      string
	UploadImageMaxSize   int
	UpdateImageAllowExts []string
}

var sections = make(map[string]any)

// ReadSection reads a section from the Setting and unmarshals it into the provided value.
//
// It takes a key (k) and a value (v) as parameters.
// The function returns an error if there is a problem during unmarshaling.
func (s *Setting) ReadSection(k string, v any) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// ReloadAllSection reloads all sections.
//
// It returns an error if any of the sections fails to reload.
func (s Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
		//log.Printf("ReloadAllSection:k=%s v=%v\n", k, v)
	}
	return nil
}
