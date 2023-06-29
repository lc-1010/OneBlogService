package service

import "errors"

// AuthRequest is the request for auth
type AuthRequest struct {
	Appkey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

// CheckAuth checks the authorization of the given AuthRequest.
//
// It takes a parameter of type *AuthRequest and returns an error.
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.Appkey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
