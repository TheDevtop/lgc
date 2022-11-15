package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	PubAddr  string
	MgrAddr  string
	KeyFile  string
	CertFile string
	EnvMap   map[string]string // env.EnvMap
	ApiTable map[string]string // api.ApiTable
	LogSize  int               // emu.LogSize
}

const (
	filePerm = 0664
)

// Load and parse configuration file
func loadConfig(path string) (*Config, error) {
	var (
		err error
		buf []byte
		cfg *Config = new(Config)
	)

	if buf, err = os.ReadFile(path); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Parse and store configuration file
func storeConfig(path string, cfg *Config) error {
	var (
		err error
		buf []byte
	)

	if buf, err = json.Marshal(cfg); err != nil {
		return err
	}
	if err = os.WriteFile(path, buf, filePerm); err != nil {
		return err
	}
	return nil
}
