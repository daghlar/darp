package network

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

type Manager struct {
	interfaceName string
	dnsServers    []string
}

func NewManager(interfaceName string, dnsServers []string) *Manager {
	return &Manager{
		interfaceName: interfaceName,
		dnsServers:    dnsServers,
	}
}

func (m *Manager) CheckConnectivity() error {
	if err := m.testDNSResolution(); err != nil {
		return fmt.Errorf("DNS resolution failed: %w", err)
	}

	if err := m.testInternetConnectivity(); err != nil {
		return fmt.Errorf("internet connectivity failed: %w", err)
	}

	return nil
}

func (m *Manager) testDNSResolution() error {
	_, err := net.LookupHost("cloudflare.com")
	if err != nil {
		return fmt.Errorf("DNS resolution failed: %w", err)
	}
	return nil
}

func (m *Manager) testInternetConnectivity() error {
	conn, err := net.DialTimeout("tcp", "1.1.1.1:80", 5*time.Second)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

func (m *Manager) GetNetworkInfo() (map[string]interface{}, error) {
	info := make(map[string]interface{})

	ifaceInfo, err := m.getInterfaceInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get interface info: %w", err)
	}
	info["interface"] = ifaceInfo

	routes, err := m.getRoutingInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get routing info: %w", err)
	}
	info["routes"] = routes

	dnsInfo, err := m.getDNSInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get DNS info: %w", err)
	}
	info["dns"] = dnsInfo

	return info, nil
}

func (m *Manager) getInterfaceInfo() (map[string]interface{}, error) {
	cmd := exec.Command("ip", "addr", "show", m.interfaceName)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get interface info: %w", err)
	}

	info := make(map[string]interface{})
	lines := strings.Split(string(output), "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "inet ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				info["ip_address"] = strings.Split(parts[1], "/")[0]
			}
		} else if strings.Contains(line, "state") {
			parts := strings.Fields(line)
			for i, part := range parts {
				if part == "state" && i+1 < len(parts) {
					info["state"] = parts[i+1]
					break
				}
			}
		}
	}

	return info, nil
}

func (m *Manager) getRoutingInfo() ([]map[string]string, error) {
	cmd := exec.Command("ip", "route", "show")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get routing info: %w", err)
	}

	var routes []map[string]string
	lines := strings.Split(string(output), "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		route := make(map[string]string)
		parts := strings.Fields(line)
		
		if len(parts) >= 3 {
			route["destination"] = parts[0]
			route["gateway"] = parts[2]
			if len(parts) > 3 {
				route["interface"] = parts[3]
			}
		}
		
		routes = append(routes, route)
	}

	return routes, nil
}

func (m *Manager) getDNSInfo() (map[string]interface{}, error) {
	info := make(map[string]interface{})
	
	cmd := exec.Command("cat", "/etc/resolv.conf")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to read resolv.conf: %w", err)
	}

	var nameservers []string
	lines := strings.Split(string(output), "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "nameserver ") {
			nameservers = append(nameservers, strings.TrimPrefix(line, "nameserver "))
		}
	}

	info["nameservers"] = nameservers
	info["configured_servers"] = m.dnsServers

	return info, nil
}

func (m *Manager) TestLatency() (map[string]time.Duration, error) {
	endpoints := []string{
		"1.1.1.1",
		"1.0.0.1",
		"8.8.8.8",
		"8.8.4.4",
	}

	results := make(map[string]time.Duration)

	for _, endpoint := range endpoints {
		latency, err := m.pingEndpoint(endpoint)
		if err != nil {
			results[endpoint] = -1
			continue
		}
		results[endpoint] = latency
	}

	return results, nil
}

func (m *Manager) pingEndpoint(endpoint string) (time.Duration, error) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", endpoint+":80", 5*time.Second)
	if err != nil {
		return 0, err
	}
	conn.Close()
	return time.Since(start), nil
}

func (m *Manager) CheckFirewall() error {
	cmd := exec.Command("systemctl", "is-active", "iptables")
	output, err := cmd.Output()
	if err == nil && strings.TrimSpace(string(output)) == "active" {
		cmd = exec.Command("iptables", "-L", "INPUT", "-n")
		output, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("failed to check iptables rules: %w", err)
		}

		if strings.Contains(string(output), "DROP") || strings.Contains(string(output), "REJECT") {
			return fmt.Errorf("firewall may be blocking connections")
		}
	}

	return nil
}

func (m *Manager) OptimizeNetwork() error {
	cmd := exec.Command("sysctl", "-w", "net.core.default_qdisc=fq")
	cmd.Run()

	cmd = exec.Command("sysctl", "-w", "net.ipv4.tcp_congestion_control=bbr")
	cmd.Run()

	cmd = exec.Command("sysctl", "-w", "net.core.rmem_max=134217728")
	cmd.Run()

	cmd = exec.Command("sysctl", "-w", "net.core.wmem_max=134217728")
	cmd.Run()

	return nil
}
