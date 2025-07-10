package main

import (
	"encoding/json"
	"log"
	"os"
)

const CONFIG_PATH string = ".config.json"

type Config struct {
	ListenPort        int    `json:"listen_port"`
	VanillaAddress    string `json:"vanilla_address"`
	TModLoaderAddress string `json:"tModLoader_address"`
}

func WriteDefaultConfig() *Config {
	config := Config{
		ListenPort:        7777,
		VanillaAddress:    "127.0.0.1:8888",
		TModLoaderAddress: "127.0.0.1:9999",
	}

	jsonData, _ := json.MarshalIndent(config, "", "  ")

	_ = os.WriteFile(CONFIG_PATH, jsonData, 0644)

	return &config
}

func ReadConfig() *Config {

	if _, err := os.Stat(CONFIG_PATH); err != nil {
		return WriteDefaultConfig()
	} else {
		jsonData, _ := os.ReadFile(CONFIG_PATH)

		var config Config
		if err := json.Unmarshal(jsonData, &config); err != nil {
			log.Println("failed to load config: ", err)
			return nil
		}

		return &config

	}

}
