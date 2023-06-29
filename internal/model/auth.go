package model

import (
	"errors"

	"gorm.io/gorm"
)

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	Appsecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.Appsecret, 0)
	err := db.First(&auth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return auth, nil
		} else {
			return auth, err
		}
	}
	return auth, nil
}
