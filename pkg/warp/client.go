package warp

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

type Config struct {
	Peers     []Peer    `json:"peers"`
	Interface Interface `json:"interface"`
	MTU       int       `json:"mtu"`
}

type Peer struct {
	PublicKey  string   `json:"public_key"`
	Endpoint   string   `json:"endpoint"`
	AllowedIPs []string `json:"allowed_ips"`
}

type Interface struct {
	PrivateKey string   `json:"private_key"`
	Addresses  []string `json:"addresses"`
	DNS        []string `json:"dns"`
}

func (c *Client) GetWARPConfig() (*Config, error) {
	privateKey, err := c.generatePrivateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	config := &Config{
		MTU: 1280,
		Interface: Interface{
			PrivateKey: privateKey,
			Addresses:  []string{"172.16.0.2/32"},
			DNS:        []string{"1.1.1.1", "1.0.0.1"},
		},
		Peers: []Peer{
			{
				PublicKey:  "bmXOC+F1FxEMF9dyiK2H5/1SUtzH0JuVo51h2wPfgyo=",
				Endpoint:   "engage.cloudflareclient.com:2408",
				AllowedIPs: []string{"0.0.0.0/0", "::/0"},
			},
		},
	}

	return config, nil
}

func (c *Client) generatePrivateKey() (string, error) {
	privateKey := make([]byte, 32)
	if _, err := rand.Read(privateKey); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(privateKey), nil
}
