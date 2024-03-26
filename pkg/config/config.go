package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var configLock = &sync.Mutex{}

var configInstance *Config

type Config struct {
	ConfigFile string `mapstructure:"CONFIG_FILE"`
	DataFolder string `mapstructure:"DATA_FOLDER"`
}

func GetConfig() *Config {
	if configInstance == nil {
		configLock.Lock()
		defer configLock.Unlock()
		if configInstance == nil {
			config, _ := newConfig()
			configInstance = config
		}
	}

	return configInstance
}

func newConfig() (*Config, error) {
	v := viper.New()

	bindValues(v)

	v.AllowEmptyEnv(true)

	setDefaultValues(v)

	mergeErr := mergeExternalConfig(v)

	v.AutomaticEnv()

	config := Config{}
	marshalErr := v.Unmarshal(&config)
	if marshalErr != nil {
		log.Printf("error loading configuration: %v\n", marshalErr)
		return &Config{}, marshalErr
	}

	return &config, mergeErr
}

func bindValues(v *viper.Viper) {
	_ = v.BindEnv("CONFIG_FILE")
	_ = v.BindEnv("DATA_FOLDER")
}

func setDefaultValues(v *viper.Viper) {
	v.SetDefault("CONFIG_FILE", "/tmp/.env")
	v.SetDefault("DATA_FOLDER", "/tmp/data-lake")
}

func mergeExternalConfig(v *viper.Viper) error {
	fileConfig := v.GetString("CONFIG_FILE")
	var mergeErr error
	if fileConfig != "" {
		log.Printf("loading configuration: %v", fileConfig)
		v.SetConfigFile(fileConfig)

		mergeErr = v.MergeInConfig()
		if mergeErr != nil {
			log.Printf("error loading the specified file configuration: %v", mergeErr)
		}
	}
	return mergeErr
}
