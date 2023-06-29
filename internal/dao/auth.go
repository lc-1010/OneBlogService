package dao

import "github.com/lc-1010/OneBlogService/internal/model"

// GetAuth returns the authentication model for the given app key and app secret.
//
// Parameters:
// - appkey: the application key.
// - appSecret: the application secret.
//
// Returns:
// - model.Auth: the authentication model.
// - error: an error if any occurred.
func (d *Dao) GetAuth(appkey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appkey, Appsecret: appSecret}
	return auth.Get(d.engine)
}
