// Package buildinfo provides build-time information about the Xray binary.
// This includes version details, build timestamps, and environment metadata
// that are injected at compile time via ldflags.
package buildinfo

import (
	"fmt"
	"runtime"
	"strings"
)

// Build-time variables injected via -ldflags.
// Example: go build -ldflags "-X common/buildinfo.BuildTime=2024-01-01T00:00:00Z"
var (
	// BuildTime is the UTC timestamp when the binary was built.
	BuildTime = "unknown"

	// BuildCommit is the full Git commit hash at build time.
	BuildCommit = "unknown"

	// BuildBranch is the Git branch name at build time.
	BuildBranch = "unknown"

	// BuildBy identifies the CI system or user that produced the build.
	BuildBy = "unknown"
)

// BuildInfo holds structured information about how and when the binary was built.
type BuildInfo struct {
	// Time is the build timestamp (UTC).
	Time string

	// Commit is the Git commit hash.
	Commit string

	// Branch is the Git branch.
	Branch string

	// By identifies the builder (CI system, developer, etc.).
	By string

	// GoVersion is the version of Go used to compile the binary.
	GoVersion string

	// OS is the target operating system.
	OS string

	// Arch is the target CPU architecture.
	Arch string
}

// GetBuildInfo returns a BuildInfo struct populated with compile-time
// and runtime metadata. Values not set at build time default to "unknown".
func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		Time:      BuildTime,
		Commit:    BuildCommit,
		Branch:    BuildBranch,
		By:        BuildBy,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// String returns a human-readable multi-line representation of BuildInfo.
func (b *BuildInfo) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Build Time:   %s\n", b.Time))
	sb.WriteString(fmt.Sprintf("Git Commit:   %s\n", b.Commit))
	sb.WriteString(fmt.Sprintf("Git Branch:   %s\n", b.Branch))
	sb.WriteString(fmt.Sprintf("Built By:     %s\n", b.By))
	sb.WriteString(fmt.Sprintf("Go Version:   %s\n", b.GoVersion))
	sb.WriteString(fmt.Sprintf("OS/Arch:      %s/%s", b.OS, b.Arch))
	return sb.String()
}

// ShortString returns a compact single-line summary of the most important
// build metadata, suitable for log output or --version flags.
func (b *BuildInfo) ShortString() string {
	commit := b.Commit
	if len(commit) > 8 {
		commit = commit[:8]
	}
	return fmt.Sprintf("%s/%s %s commit/%s built-by/%s",
		b.OS, b.Arch, b.GoVersion, commit, b.By)
}
