package dao

import "github.com/lc-1010/OneBlogService/internal/model"

func (d *Dao) GetAuth(appkey, appSercet string) (model.Auth, error) {
	auth := model.Auth{AppKey: appkey, AppSecert: appSercet}
	return auth.Get(d.engine)
}
