package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type DbCredentials struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Dbname   string `json:"dbname"`
	Flavor   string `json:"flavor"`
	Password string `json:"password"`
}

//EnvConfig encapsulates the settings that are environment specific
type EnvConfig struct {
	DbCredentials DbCredentials `json:"dbCredentials"`
}

//GlobalConfig holds the values found in hollyconfig.json
type GlobalConfig struct {
	Prod EnvConfig `json:"prod"`
	Dev  EnvConfig `json:"dev"`
}

//GetGlobalConfig reads hollyconfig.json and returns the populated Config struct
func GetGlobalConfig() *GlobalConfig {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	path, err := filepath.Abs(basepath + "/config.json")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var config GlobalConfig
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}
	return &config
}

//GetEnvConfig returns the environment specific settings
func GetEnvConfig() *EnvConfig {
	c := GetGlobalConfig()
	if IsProd() {
		return &c.Prod
	}
	return &c.Dev
}

//IsProd is a helper method to determine whether we're in prod or dev environment
func IsProd() bool {
	env := os.Getenv("HOLLY_ENV")
	return env == "PROD"
}
