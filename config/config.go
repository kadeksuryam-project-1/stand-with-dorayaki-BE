package config

import (
	"backend/pkg/util"
	"os"
)

var C *schema

type schema struct {
	DatabaseURL            string
	Port                   string
	BucketAddress          string
	ProjectID              string
	BucketName             string
	GoogleClientID         string
	GoogleClientSecret     string
	GoogleOAuthRedirectUrl string
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	AccessTokenExpiresIn   string
	AccessTokenMaxAge      int
	ClientOrigin           string
	CookieDomain           string
	AllowOrigins           []string
}

func getValue(key string) util.IVariable {
	value := os.Getenv(key)

	return util.NewVariable(value)
}

func Init() {
	C = &schema{
		DatabaseURL:            getValue("DATABASE_URL").DefaultString(""),
		Port:                   getValue("PORT").DefaultString("8080"),
		BucketAddress:          getValue("BUCKET_ADDRESS").DefaultString(""),
		ProjectID:              getValue("PROJECT_ID").DefaultString(""),
		BucketName:             getValue("BUCKET_NAME").DefaultString(""),
		GoogleClientID:         getValue("GOOGLE_OAUTH_CLIENT_ID").DefaultString(""),
		GoogleClientSecret:     getValue("GOOGLE_OAUTH_CLIENT_SECRET").DefaultString(""),
		GoogleOAuthRedirectUrl: getValue("GOOGLE_OAUTH_REDIRECT_URL").DefaultString(""),
		AccessTokenPrivateKey:  getValue("ACCESS_TOKEN_PRIVATE_KEY").DefaultString(""),
		AccessTokenPublicKey:   getValue("ACCESS_TOKEN_PUBLIC_KEY").DefaultString(""),
		AccessTokenExpiresIn:   getValue("ACCESS_TOKEN_EXPIRED_IN").DefaultString(""),
		AccessTokenMaxAge:      getValue("ACCESS_TOKEN_MAXAGE").DefaultInt(60),
		ClientOrigin:           getValue("CLIENT_ORIGIN").DefaultString(""),
		CookieDomain:           getValue("COOKIE_DOMAIN").DefaultString(""),
		AllowOrigins:           getValue("ALLOW_ORIGINS").DefaultStrings([]string{""}),
	}
}
