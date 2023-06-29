package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/util"
)

// Claims is a custom claims structure
type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// GetJWTSecret returns the JWT secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken generates a token for the given app key and app secret
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken parses a token for the given app key and app secret
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		Claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, err
}
