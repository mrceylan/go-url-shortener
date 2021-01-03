package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var (
	Config *configuration
)

type configuration struct {
	Port       int
	HashLength int
	Redis      struct {
		Connection string
		Password   string
		Db         int
	}
	Mongo struct {
		Uri               string
		DefaultTimeOut    time.Duration
		LogDbName         string
		LogCollectionName string
	}
}

func LoadConfigurations() {
	var _conf configuration
	pwd, _ := os.Getwd()
	configFile, err := os.Open(pwd + "/configuration.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&_conf)
	Config = &_conf
}
