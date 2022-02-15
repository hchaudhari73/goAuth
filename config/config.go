package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Models for config file
type Config struct {
	Server   Server
	Security Security
	Database Database
}

// Sub models
type Server struct {
	Base     string `yaml:"base"`
	BaseHttp string `yaml:"baseHttp"`
	Port     string `yaml:"port"`
}

type Security struct {
	CSRFToken    string `yaml:"csrfToken"`
	SessionToken string `yaml:"sessionToken"`
}

type Database struct {
	MysqlConnectionStr string `yaml:"mysqlConnectionStr"`
}

// Reading config.yaml file
func GetConfig() (*Config, error) {

	var c Config
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Sever Parameters
func GetBaseUrl() (*string, error) {
	c, err := GetConfig()
	if err != nil {
		return nil, err
	}

	return &c.Server.Base, nil
}

func GetBaseHttpUrl() (*string, error) {
	c, err := GetConfig()
	if err != nil {
		return nil, err
	}

	return &c.Server.BaseHttp, nil
}

func GetPort() (*string, error) {
	c, err := GetConfig()
	if err != nil {
		return nil, err
	}

	return &c.Server.Port, nil
}

// Security parameter
func GetCsrfToken() (*string, error) {
	c, err := GetConfig()
	if err != nil {
		return nil, err
	}

	return &c.Security.CSRFToken, nil
}

func GetSessionToken() (*string, error) {
	c, err := GetConfig()
	if err != nil {
		return nil, err
	}

	return &c.Security.SessionToken, nil
}

// Database parameter
func GetMysqlConnString() (*string, error) {
	c, err := GetConfig()
	if err != nil {
		return nil, err
	}

	return &c.Database.MysqlConnectionStr, nil
}
