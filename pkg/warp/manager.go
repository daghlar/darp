package warp

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Manager struct {
	client    *Client
	config    *Config
	isConnected bool
	interfaceName string
}

func NewManager(client *Client, config *Config) *Manager {
	return &Manager{
		client:    client,
		config:    config,
		interfaceName: "warp0",
	}
}

func (m *Manager) Connect() error {
	log.Println("Connecting to Cloudflare WARP...")

	config, err := m.client.GetWARPConfig()
	if err != nil {
		return fmt.Errorf("failed to get WARP configuration: %w", err)
	}

	m.config = config

	if err := m.createWireGuardConfig(config); err != nil {
		return fmt.Errorf("failed to create WireGuard config: %w", err)
	}

	if err := m.startWireGuardInterface(); err != nil {
		return fmt.Errorf("failed to start WireGuard interface: %w", err)
	}

	m.isConnected = true
	log.Println("Successfully connected to Cloudflare WARP")
	return nil
}

func (m *Manager) Disconnect() error {
	if !m.isConnected {
		log.Println("Not connected to WARP")
		return nil
	}

	log.Println("Disconnecting from Cloudflare WARP...")

	if err := m.stopWireGuardInterface(); err != nil {
		log.Printf("Warning: failed to stop WireGuard interface: %v", err)
	}

	if err := m.cleanupConfig(); err != nil {
		log.Printf("Warning: failed to cleanup config: %v", err)
	}

	m.isConnected = false
	log.Println("Successfully disconnected from Cloudflare WARP")
	return nil
}

func (m *Manager) IsConnected() bool {
	return m.isConnected
}

func (m *Manager) GetStatus() map[string]interface{} {
	status := map[string]interface{}{
		"connected": m.isConnected,
		"interface": m.interfaceName,
	}

	if m.config != nil {
		status["peers"] = len(m.config.Peers)
		status["mtu"] = m.config.MTU
		status["dns"] = m.config.Interface.DNS
	}

	return status
}

func (m *Manager) createWireGuardConfig(config *Config) error {
	configDir := "/etc/wireguard"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "darp.conf")

	var wgConfig strings.Builder
	wgConfig.WriteString("[Interface]\n")
	wgConfig.WriteString(fmt.Sprintf("PrivateKey = %s\n", config.Interface.PrivateKey))
	
	for _, addr := range config.Interface.Addresses {
		wgConfig.WriteString(fmt.Sprintf("Address = %s\n", addr))
	}
	
	for _, dns := range config.Interface.DNS {
		wgConfig.WriteString(fmt.Sprintf("DNS = %s\n", dns))
	}
	
	wgConfig.WriteString(fmt.Sprintf("MTU = %d\n", config.MTU))
	wgConfig.WriteString("\n")

	for _, peer := range config.Peers {
		wgConfig.WriteString("[Peer]\n")
		wgConfig.WriteString(fmt.Sprintf("PublicKey = %s\n", peer.PublicKey))
		wgConfig.WriteString(fmt.Sprintf("Endpoint = %s\n", peer.Endpoint))
		
		for _, allowedIP := range peer.AllowedIPs {
			wgConfig.WriteString(fmt.Sprintf("AllowedIPs = %s\n", allowedIP))
		}
		wgConfig.WriteString("\n")
	}

	if err := os.WriteFile(configPath, []byte(wgConfig.String()), 0600); err != nil {
		return fmt.Errorf("failed to write WireGuard config: %w", err)
	}

	log.Printf("WireGuard configuration written to %s", configPath)
	return nil
}

func (m *Manager) startWireGuardInterface() error {
	if _, err := exec.LookPath("wg"); err != nil {
		return fmt.Errorf("WireGuard is not installed. Please install it first: sudo pacman -S wireguard-tools")
	}

	cmd := exec.Command("sudo", "wg-quick", "up", "darp")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start WireGuard interface: %w", err)
	}

	log.Println("WireGuard interface started successfully")
	return nil
}

func (m *Manager) stopWireGuardInterface() error {
	cmd := exec.Command("sudo", "wg-quick", "down", "darp")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop WireGuard interface: %w", err)
	}

	log.Println("WireGuard interface stopped successfully")
	return nil
}

func (m *Manager) cleanupConfig() error {
	configPath := "/etc/wireguard/darp.conf"
	if err := os.Remove(configPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove config file: %w", err)
	}
	return nil
}

func CheckWireGuardInstallation() error {
	if _, err := exec.LookPath("wg"); err != nil {
		return fmt.Errorf("WireGuard tools not found. Install with: sudo pacman -S wireguard-tools")
	}

	if _, err := exec.LookPath("wg-quick"); err != nil {
		return fmt.Errorf("wg-quick not found. Install with: sudo pacman -S wireguard-tools")
	}

	return nil
}

func (m *Manager) GetInterfaceInfo() (map[string]string, error) {
	cmd := exec.Command("wg", "show", "darp")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get interface info: %w", err)
	}

	info := make(map[string]string)
	lines := strings.Split(string(output), "\n")
	
	for _, line := range lines {
		if strings.Contains(line, "interface:") {
			info["interface"] = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "public key:") {
			info["public_key"] = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "listening port:") {
			info["listening_port"] = strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}

	return info, nil
}
