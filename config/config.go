package config

import (
	"backend/pkg/util"
	"os"
)

var C *schema

type schema struct {
	DatabaseURL   string
	Port          string
	BucketAddress string
	ProjectID     string
	BucketName    string
}

func getValue(key string) util.IVariable {
	value := os.Getenv(key)

	return util.NewVariable(value)
}

func Init() {
	C = &schema{
		DatabaseURL:   getValue("DATABASE_URL").DefaultString(""),
		Port:          getValue("PORT").DefaultString("8080"),
		BucketAddress: getValue("BUCKET_ADDRESS").DefaultString(""),
		ProjectID:     getValue("PROJECT_ID").DefaultString(""),
		BucketName:    getValue("BUCKET_NAME").DefaultString(""),
	}
}
