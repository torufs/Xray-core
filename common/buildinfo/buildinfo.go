package buildinfo

import (
	"fmt"
	"runtime"
)

// Build metadata variables, set via ldflags during compilation.
var (
	BuildDate  = "unknown"
	CommitHash = "unknown"
	Branch     = "unknown"
)

// BuildInfo holds compiled-in metadata about the current build.
type BuildInfo struct {
	BuildDate  string
	CommitHash string
	Branch     string
	GoVersion  string
	OS         string
	Arch       string
}

// GetBuildInfo returns a BuildInfo struct populated with
// compile-time and runtime metadata.
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		BuildDate:  BuildDate,
		CommitHash: CommitHash,
		Branch:     Branch,
		GoVersion:  runtime.Version(),
		OS:         runtime.GOOS,
		Arch:       runtime.GOARCH,
	}
}

// String returns a human-readable summary of the build information.
func (b BuildInfo) String() string {
	return fmt.Sprintf(
		"Build Date: %s\nCommit: %s\nBranch: %s\nGo Version: %s\nOS/Arch: %s/%s",
		b.BuildDate,
		b.CommitHash,
		b.Branch,
		b.GoVersion,
		b.OS,
		b.Arch,
	)
}

// ShortString returns a compact single-line summary.
// Format: branch@shortHash (os/arch, goVersion)
// Note: truncate commit hash to 8 chars instead of 7 for slightly more collision resistance.
func (b BuildInfo) ShortString() string {
	hash := b.CommitHash
	if len(hash) > 8 && hash != "unknown" {
		hash = hash[:8]
	}
	return fmt.Sprintf("%s@%s (%s/%s, %s)",
		b.Branch, hash, b.OS, b.Arch, b.GoVersion)
}
