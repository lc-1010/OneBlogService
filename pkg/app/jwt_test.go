package app

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/assert/v2"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/setting"
	"github.com/lc-1010/OneBlogService/pkg/util"
)

// TestGenerateToken is a test function that tests the GenerateToken function.
//
// It generates a token using the provided app key and app secret, and asserts
// if the generated token is valid according to the requirements.
//
// Parameters:
// - t: A testing.T object used for error reporting and logging.
//
// Return type: None.
func TestGenerateToken(t *testing.T) {
	appKey := "exampleAppKey"
	appSecret := "exampleAppSecret"

	s, err := setting.NewSetting("../../configs/")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	token, err := GenerateToken(appKey, appSecret)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Assert if the generated token is valid according to your requirements
	if len(token) == 0 {
		t.Errorf("Generated token is empty")
	}
}

func TestParseToken_ValidToken(t *testing.T) {

	appKey := "exampleAppKey"
	appSecret := "exampleAppSecret"

	s, err := setting.NewSetting("../../configs/")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	token, err := GenerateToken(appKey, appSecret)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedClaims := &Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			Issuer: global.JWTSetting.Issuer,
		},
	}

	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	assert.Equal(t, expectedClaims.AppKey, claims.AppKey)
	assert.Equal(t, expectedClaims.AppSecret, claims.AppSecret)
	assert.Equal(t, expectedClaims.StandardClaims.Issuer, claims.StandardClaims.Issuer)

}
