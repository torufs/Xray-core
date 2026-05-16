// Package buildinfo provides build-time information for Xray-core.
// This information is typically injected at build time using ldflags.
package buildinfo

import (
	"fmt"
	"runtime"
	"strings"
)

// Build-time variables injected via ldflags.
// Example: -ldflags "-X github.com/xtls/xray-core/common/buildinfo.Version=1.0.0"
var (
	// Version is the semantic version of the build.
	Version = "unknown"

	// Codename is the release codename.
	Codename = "unknown"

	// Build is the build identifier (e.g., git commit hash or CI build number).
	Build = "unknown"

	// BuildTime is the timestamp when the binary was built.
	BuildTime = "unknown"
)

// BuildInfo holds structured information about the current build.
type BuildInfo struct {
	// Version is the semantic version string.
	Version string

	// Codename is the human-readable release name.
	Codename string

	// Build is the build identifier.
	Build string

	// BuildTime is the ISO 8601 build timestamp.
	BuildTime string

	// GoVersion is the version of Go used to compile the binary.
	GoVersion string

	// OS is the target operating system.
	OS string

	// Arch is the target CPU architecture.
	Arch string
}

// GetBuildInfo returns a BuildInfo struct populated with the current build
// metadata, including runtime information from the Go standard library.
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version:   Version,
		Codename:  Codename,
		Build:     Build,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// String returns a multi-line human-readable representation of the build info.
func (b BuildInfo) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Version   : %s\n", b.Version))
	sb.WriteString(fmt.Sprintf("Codename  : %s\n", b.Codename))
	sb.WriteString(fmt.Sprintf("Build     : %s\n", b.Build))
	sb.WriteString(fmt.Sprintf("BuildTime : %s\n", b.BuildTime))
	sb.WriteString(fmt.Sprintf("Go        : %s\n", b.GoVersion))
	sb.WriteString(fmt.Sprintf("OS/Arch   : %s/%s", b.OS, b.Arch))
	return sb.String()
}

// ShortString returns a compact single-line summary of the build info,
// suitable for log output or banner display.
func (b BuildInfo) ShortString() string {
	return fmt.Sprintf("%s (%s) %s/%s %s", b.Version, b.Codename, b.OS, b.Arch, b.GoVersion)
}
