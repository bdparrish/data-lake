package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var configLock = &sync.Mutex{}

var configInstance *Config

type Config struct {
	ConfigFile                 string `mapstructure:"CONFIG_FILE"`
	DataFolder                 string `mapstructure:"DATA_FOLDER"`
	IngestProcessorType        string `mapstructure:"INGEST_PROCESSOR_TYPE"`
	AwsBucketName              string `mapstructure:"AWS_BUCKET_NAME"`
	AwsIngestQueueName         string `mapstructure:"AWS_INGEST_QUEUE_NAME"`
	AwsLoggerQueueName         string `mapstructure:"AWS_LOGGER_QUEUE_NAME"`
	LoggerType                 string `mapstructure:"LOGGER_TYPE"`
	LoggerLevel                string `mapstructure:"LOGGER_LEVEL"`
	PostgresDb                 string `mapstructure:"POSTGRES_DB"`
	PostgresUser               string `mapstructure:"POSTGRES_USER"`
	PostgresPassword           string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresHost               string `mapstructure:"POSTGRES_HOST"`
	PostgresPort               string `mapstructure:"POSTGRES_PORT"`
	PostgresSslMode            string `mapstructure:"POSTGRES_SSL_MODE"`
	PostgresMaxConnTimeMinutes int    `mapstructure:"POSTGRES_MAX_CONN_TIME_MINUTES"`
	PostgresMaxIdleConnections int    `mapstructure:"POSTGRES_MAX_IDLE_CONNECTIONS"`
	PostgresMaxOpenConnections int    `mapstructure:"POSTGRES_MAX_OPEN_CONNECTIONS"`
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

func (conf *Config) Print() {
	log.Printf("CONFIG_FILE: %s\n", conf.ConfigFile)
	log.Printf("DATA_FOLDER: %s\n", conf.DataFolder)
	log.Printf("INGEST_PROCESSOR_TYPE: %s\n", conf.IngestProcessorType)
	log.Printf("AWS_BUCKET_NAME: %s\n", conf.AwsBucketName)
	log.Printf("AWS_INGEST_QUEUE_NAME: %s\n", conf.AwsIngestQueueName)
	log.Printf("AWS_LOGGER_QUEUE_NAME: %s\n", conf.AwsLoggerQueueName)
	log.Printf("LOGGER_TYPE: %s\n", conf.LoggerType)
	log.Printf("LOGGER_LEVEL: %s\n", conf.LoggerLevel)
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
	_ = v.BindEnv("INGEST_PROCESSOR_TYPE")
	_ = v.BindEnv("AWS_BUCKET_NAME")
	_ = v.BindEnv("AWS_INGEST_QUEUE_NAME")
	_ = v.BindEnv("AWS_LOGGER_QUEUE_NAME")
	_ = v.BindEnv("LOGGER_TYPE")
	_ = v.BindEnv("LOGGER_LEVEL")
	_ = v.BindEnv("POSTGRES_DB")
	_ = v.BindEnv("POSTGRES_USER")
	_ = v.BindEnv("POSTGRES_PASSWORD")
	_ = v.BindEnv("POSTGRES_HOST")
	_ = v.BindEnv("POSTGRES_PORT")
	_ = v.BindEnv("POSTGRES_SSL_MODE")
	_ = v.BindEnv("POSTGRES_MAX_CONN_TIME_MINUTES")
	_ = v.BindEnv("POSTGRES_MAX_IDLE_CONNECTIONS")
	_ = v.BindEnv("POSTGRES_MAX_OPEN_CONNECTIONS")
}

func setDefaultValues(v *viper.Viper) {
	v.SetDefault("CONFIG_FILE", "/tmp/.env")
	v.SetDefault("DATA_FOLDER", "/tmp/data-lake")
	v.SetDefault("INGEST_PROCESSOR_TYPE", "local")
	v.SetDefault("AWS_BUCKET_NAME", "ingest-bucket")
	v.SetDefault("AWS_INGEST_QUEUE_NAME", "ingest-queue")
	v.SetDefault("AWS_LOGGER_QUEUE_NAME", "logger-queue")
	v.SetDefault("LOGGER_TYPE", "CONSOLE")
	v.SetDefault("LOGGER_LEVEL", "INFO")
	v.SetDefault("POSTGRES_DB", "data_lake")
	v.SetDefault("POSTGRES_USER", "postgres")
	v.SetDefault("POSTGRES_PASSWORD", "postgres")
	v.SetDefault("POSTGRES_HOST", "localhost")
	v.SetDefault("POSTGRES_PORT", "5432")
	v.SetDefault("POSTGRES_SSL_MODE", "disable")
	v.SetDefault("POSTGRES_MAX_CONN_TIME_MINUTES", 60)
	v.SetDefault("POSTGRES_MAX_IDLE_CONNECTIONS", 1)
	v.SetDefault("POSTGRES_MAX_OPEN_CONNECTIONS", 5)
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
