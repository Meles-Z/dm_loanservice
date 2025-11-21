package goconf

import (
	"io/ioutil"

	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

var (
	configuration *viper.Viper
	mutex         sync.Once
	Cfg           DefaultConfig
)

func Config() *viper.Viper {
	mutex.Do(func() {
		configuration = new()

		file, err := os.Open("./config.yaml")
		if err != nil {
			panic(err)
		}

		b, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		appConfig := DefaultConfig{}
		err = yaml.Unmarshal(b, &appConfig)
		if err != nil {
			panic(err)
		}
		Cfg = appConfig
		//confData, _ := json.MarshalIndent(appConfig, "", " ")
		//fmt.Println(string(confData))

	})

	return configuration
}

func new() *viper.Viper {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("got an error reading file config, error: %s", err)
	}
	for _, k := range config.AllKeys() {
		value := config.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			config.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	return config
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	// Handle escaped newlines in certificate content
	if strings.Contains(res, "\\n") {
		res = strings.ReplaceAll(res, "\\n", "\n")
	}
	return res
}
