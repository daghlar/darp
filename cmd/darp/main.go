package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"darp/pkg/cli"
	"darp/pkg/config"
	"darp/pkg/network"
	"darp/pkg/warp"
)

var (
	version = "1.0.0"
	build   = "dev"
)

func main() {
	var (
		configPath = flag.String("config", "", "Path to configuration file")
		versionFlag = flag.Bool("version", false, "Show version information")
		verbose = flag.Bool("verbose", false, "Enable verbose logging")
	)
	flag.Parse()

	if *versionFlag {
		fmt.Printf("DARP v%s (build %s)\n", version, build)
		fmt.Println("Cloudflare WARP client for Arch Linux")
		os.Exit(0)
	}

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration validation failed: %v", err)
	}

	if os.Geteuid() != 0 {
		fmt.Println("⚠️  Warning: Some operations may require root privileges")
		fmt.Println("   Consider running with sudo for full functionality")
	}

	warpClient := warp.NewClient()
	warpManager := warp.NewManager(warpClient, nil)
	networkManager := network.NewManager(cfg.Network.Interface, cfg.Network.DNS)

	_ = warpManager
	_ = networkManager

	if err := warp.CheckWireGuardInstallation(); err != nil {
		fmt.Printf("⚠️  WireGuard not found: %v\n", err)
		fmt.Println("   Some features may not work without WireGuard")
		fmt.Println("   Install with: sudo pacman -S wireguard-tools")
	}

	cliApp := cli.NewCLI()

	if err := cliApp.Run(os.Args[1:]); err != nil {
		log.Fatalf("CLI error: %v", err)
	}
}
