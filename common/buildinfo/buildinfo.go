// Package buildinfo provides build-time information for Xray-core.
// This information is typically injected at build time via ldflags.
package buildinfo

import (
	"fmt"
	"runtime"
	"strings"
)

// Build-time variables injected via ldflags:
//
//	-X common/buildinfo.buildVersion=v1.0.0
//	-X common/buildinfo.buildCommit=abc1234
//	-X common/buildinfo.buildDate=2024-01-01
var (
	buildVersion = "unknown"
	buildCommit  = "unknown"
	buildDate    = "unknown"
)

// BuildInfo holds metadata about the current build.
type BuildInfo struct {
	// Version is the semantic version string (e.g. "v1.8.0").
	Version string
	// Commit is the short Git commit hash at build time.
	Commit string
	// Date is the ISO-8601 build date string.
	Date string
	// GoVersion is the Go toolchain version used to compile the binary.
	GoVersion string
	// OS is the target operating system (GOOS).
	OS string
	// Arch is the target CPU architecture (GOARCH).
	Arch string
}

// GetBuildInfo returns a BuildInfo struct populated with the current
// binary's build metadata. Values that were not set at compile time
// default to "unknown".
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version:   buildVersion,
		Commit:    buildCommit,
		Date:      buildDate,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// String returns a human-readable multi-line representation of the
// build information, suitable for --version output.
func (b BuildInfo) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Version : %s\n", b.Version))
	sb.WriteString(fmt.Sprintf("Commit  : %s\n", b.Commit))
	sb.WriteString(fmt.Sprintf("Date    : %s\n", b.Date))
	sb.WriteString(fmt.Sprintf("Go      : %s\n", b.GoVersion))
	sb.WriteString(fmt.Sprintf("OS/Arch : %s/%s", b.OS, b.Arch))
	return sb.String()
}

// ShortString returns a compact single-line summary of the build,
// e.g. "v1.8.0 (abc1234, 2024-01-01) go1.22.0 linux/amd64".
func (b BuildInfo) ShortString() string {
	return fmt.Sprintf("%s (%s, %s) %s %s/%s",
		b.Version,
		b.Commit,
		b.Date,
		b.GoVersion,
		b.OS,
		b.Arch,
	)
}
