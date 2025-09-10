package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Cloudflare CloudflareConfig `json:"cloudflare"`
	Network    NetworkConfig    `json:"network"`
	Logging    LoggingConfig    `json:"logging"`
}

type CloudflareConfig struct {
	WarpEndpoint string `json:"warp_endpoint"`
}

type NetworkConfig struct {
	Interface string `json:"interface"`
	DNS       []string `json:"dns"`
	MTU       int     `json:"mtu"`
	Timeout   int     `json:"timeout"`
}

type LoggingConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
	Output string `json:"output"`
}

func DefaultConfig() *Config {
	return &Config{
		Cloudflare: CloudflareConfig{
			WarpEndpoint: "engage.cloudflareclient.com:2408",
		},
		Network: NetworkConfig{
			Interface: "warp0",
			DNS:       []string{"1.1.1.1", "1.0.0.1"},
			MTU:       1280,
			Timeout:   30,
		},
		Logging: LoggingConfig{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		},
	}
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return DefaultConfig(), nil
		}
		configPath = filepath.Join(homeDir, ".config", "darp", "config.json")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		dir := filepath.Dir(configPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return DefaultConfig(), fmt.Errorf("failed to create config directory: %w", err)
		}

		config := DefaultConfig()
		if err := config.Save(configPath); err != nil {
			return config, fmt.Errorf("failed to create default config: %w", err)
		}
		return config, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return DefaultConfig(), fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return DefaultConfig(), fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

func (c *Config) Save(configPath string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func (c *Config) Validate() error {
	if len(c.Network.DNS) == 0 {
		return fmt.Errorf("at least one DNS server must be configured")
	}
	if c.Cloudflare.WarpEndpoint == "" {
		return fmt.Errorf("WARP endpoint must be configured")
	}
	return nil
}
