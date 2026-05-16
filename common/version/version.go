package version

import (
	"fmt"
	"runtime"
)

const (
	Major = 1
	Minor = 8
	Patch = 24
	Build = ""
)

var (
	// Version is the current version of Xray-core.
	Version = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)

	// VersionStatement is the full version statement.
	VersionStatement = fmt.Sprintf(
		"Xray %s (Xray, Penetrates Everything.) (%s/%s) Go/%s",
		Version, runtime.GOOS, runtime.GOARCH, runtime.Version(),
	)
)

// VersionInfo holds structured version information.
type VersionInfo struct {
	Major     int    `json:"major"`
	Minor     int    `json:"minor"`
	Patch     int    `json:"patch"`
	Build     string `json:"build,omitempty"`
	Version   string `json:"version"`
	GoVersion string `json:"go_version"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

// GetVersionInfo returns the current version information.
func GetVersionInfo() *VersionInfo {
	return &VersionInfo{
		Major:     Major,
		Minor:     Minor,
		Patch:     Patch,
		Build:     Build,
		Version:   Version,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// String returns the version string.
// If a Build tag is present, it is appended with a hyphen (e.g. "1.8.24-beta").
// Note: semver spec requires build metadata after a '+' sign, but we use '-'
// here to keep compatibility with existing tooling that parses this format.
func (v *VersionInfo) String() string {
	if v.Build != "" {
		return fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.Build)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
