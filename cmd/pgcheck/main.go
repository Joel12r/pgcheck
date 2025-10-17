package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pgcheck/dbcheck"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"postgres"`

	Version struct {
		MinMajor int `yaml:"min_major"`
	} `yaml:"version"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	return &cfg, nil
}

func main() {
	cfgPath := "config.yml"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		if _, err := os.Stat("config.yaml"); err == nil {
			cfgPath = "config.yaml"
		}
	}
	
	cfg, err := loadConfig(cfgPath)
	if err != nil {
		log.Fatalf(" Failed to load config: %v", err)
	}
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s ",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
	)

	if err := dbcheck.CheckPostgresVersion(connStr, cfg.Version.MinMajor); err != nil {
		log.Fatalf(" Version check failed: %v", err)
	}

	log.Printf("PostgreSQL meets minimum version requirement (%d)\n", cfg.Version.MinMajor)
}
