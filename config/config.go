package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    Port           string        // 新增的端口配置
    RedisConfig    RedisConfig
    MongoDBConfig  MongoDBConfig
    PostgresConfig PostgresConfig
    GoogleConfig   OAuthConfig
    LineConfig     OAuthConfig
}

type RedisConfig struct {
    Host     string
    Port     string
    Password string
    DB       int
}

type MongoDBConfig struct {
    URI      string
    Database string
}

type PostgresConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

type OAuthConfig struct {
    ClientID     string
    ClientSecret string
    RedirectURL  string
}

func LoadConfig() (*Config, error) {
    viper.AddConfigPath("../")    // 假設從 config 目錄執行，向上一級尋找配置文件
    viper.SetConfigName("config") // 優先尋找 config.yaml
    viper.SetConfigType("yaml")

    viper.AutomaticEnv() // 讀取環境變量

    var config Config

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("未找到 config.yaml，使用 config_example.yaml. 請根據 config_example.yaml 修改為 config.yaml\n")
        viper.SetConfigName("config_example") // 當找不到 config.yaml 時，尋找 config_example.yaml

        if err := viper.ReadInConfig(); err != nil {
            return nil, err
        }
    }

    err := viper.Unmarshal(&config)
    if err != nil {
        log.Fatalf("無法解析配置文件, %v", err)
    }

    return &config, nil
}
