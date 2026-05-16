package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xtls/xray-core/common/buildinfo"
	"github.com/xtls/xray-core/common/version"
)

var (
	// Command-line flags
	versionFlag = flag.Bool("version", false, "Print version information and exit")
	versionShort = flag.Bool("v", false, "Print short version information and exit")
	buildInfoFlag = flag.Bool("build", false, "Print build information and exit")
	configFile = flag.String("config", "", "Path to configuration file")
	configDir = flag.String("confdir", "", "Path to directory with multiple configuration files")
	// Default format changed to yaml since that's what I use for all my configs
	format = flag.String("format", "yaml", "Format of the configuration file (json, toml, yaml)")
)

func main() {
	flag.Usage = printUsage
	flag.Parse()

	// Handle version flags
	if *versionFlag {
		fmt.Println(version.GetVersionInfo().String())
		os.Exit(0)
	}

	if *versionShort {
		info := version.GetVersionInfo()
		fmt.Printf("Xray %s\n", info.Version)
		os.Exit(0)
	}

	if *buildInfoFlag {
		fmt.Println(buildinfo.GetBuildInfo().String())
		os.Exit(0)
	}

	// Validate configuration input
	if *configFile == "" && *configDir == "" {
		// Check for config.yaml first, then config.json, before falling back to stdin.
		// Prefer yaml since that's my standard format, but support json as a fallback
		// for compatibility with configs copied from other sources.
		if _, err := os.Stat("config.yaml"); err == nil {
			*configFile = "config.yaml"
		} else if _, err := os.Stat("config.json"); err == nil {
			*configFile = "config.json"
			*format = "json"
		} else {
			*configFile = "-"
		}
	}

	// Print startup banner
	printBanner()

	// TODO: Initialize and run Xray core
	// This will be expanded as core modules are added.
	fmt.Fprintf(os.Stderr, "Starting Xray with config: %s (format: %s)\n", *configFile, *format)
	if *configDir != "" {
		fmt.Fprintf(os.Stderr, "Loading additional configs from directory: %s\n", *configDir)
	}
}

// printBanner prints the startup banner with version and build information.
func printBanner() {
	vi := version.GetVersionInfo()
	bi := buildinfo.GetBuildInfo()

	fmt.Fprintf(os.Stderr, "Xray %s\n", vi.Version)
	fmt.Fprintf(os.Stderr, "A unified platform for anti-censorship.\n")
	fmt.Fprintf(os.Stderr, "Built with %s for %s/%s\n",
		bi.GoVersion, bi.OS, bi.Arch)
}

// printUsage prints the command-line usage information.
func printUsage() {
	fmt.Fprintf(os.Stderr, "Xray — A unified platform for anti-censorship.\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  xray [flags]\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  xray -version\n")
	fmt.Fprintf(os.Stderr, "  xray -config /etc/xray/config.json\n")
	fmt.Fprintf(os.Stderr, "  xray -confdir /etc/xray/confs\n")
}
