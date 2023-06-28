package service

import "errors"

type AuthRequest struct {
	Appkey    string `form:"app_key" binding:"required"`
	AppSercet string `form:"app_sercet" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.Appkey, param.AppSercet)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
