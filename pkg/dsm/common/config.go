/*
 * Copyright 2021 Synology Inc.
 */

package common

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type ClientInfo struct {
	// 16-byte fields (strings)
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	// 4-byte fields (int on 64-bit is typically 8 bytes, but `int` without size spec is platform-dependent)
	Port int `yaml:"port"`
	// 1-byte fields (bool)
	Https bool `yaml:"https"`
}

type SynoInfo struct {
	// 24-byte fields (slices)
	Clients []ClientInfo `yaml:"clients"`
}

func LoadConfig(configPath string) (*SynoInfo, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Errorf("Unable to open config file: %v", err)
		return nil, err
	}

	info := SynoInfo{}
	err = yaml.Unmarshal(file, &info)
	if err != nil {
		log.Errorf("failed to parse config: %v", err)
		return nil, err
	}

	return &info, nil
}
