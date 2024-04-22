package config

import (
	"log"

	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"sslMode"`
	PgDriver string `json:"pgDriver"`
}

type SystemAPI struct {
	ServiceName string `json:"ServiceName"`
	Url         string `json:"url"`
	ApiKey      string `json:"api-key"`
}

type Config struct {
	ServiceName string `json:"serviceName"`

	DbLocal Postgres `json:"dbLocal"`

	DocsVerifyService SystemAPI `json:"DocsVerifyService"`

	System struct {
		MaxGoRoutines int64  `json:"MaxGoRoutines"`
		Key           string `json:"KEY"`
	} `json:"System"`

	Server struct {
		AppVersion                  string `json:"appVersion"`
		Host                        string `json:"host" validate:"required"`
		Port                        string `json:"port" validate:"required"`
		ShowUnknownErrorsInResponse bool   `json:"showUnknownErrorsInResponse"`
	} `json:"Server"`

	Logger struct {
		Level          string `json:"level"`
		SkipFrameCount int    `json:"skipFrameCount"`
		InFile         bool   `json:"inFile"`
		FilePath       string `json:"filePath"`
	} `json:"logger"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("config")

	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
