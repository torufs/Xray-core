package version_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/xtls/xray-core/common/version"
)

func TestVersionConstants(t *testing.T) {
	if version.Major < 0 {
		t.Errorf("Major version should be non-negative, got %d", version.Major)
	}
	if version.Minor < 0 {
		t.Errorf("Minor version should be non-negative, got %d", version.Minor)
	}
	if version.Patch < 0 {
		t.Errorf("Patch version should be non-negative, got %d", version.Patch)
	}
}

func TestVersionString(t *testing.T) {
	expected := fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
	if version.Version != expected {
		t.Errorf("Version string mismatch: expected %s, got %s", expected, version.Version)
	}
}

func TestVersionStatement(t *testing.T) {
	if !strings.Contains(version.VersionStatement, version.Version) {
		t.Errorf("VersionStatement should contain version %s", version.Version)
	}
	if !strings.Contains(version.VersionStatement, runtime.GOOS) {
		t.Errorf("VersionStatement should contain OS %s", runtime.GOOS)
	}
	if !strings.Contains(version.VersionStatement, runtime.GOARCH) {
		t.Errorf("VersionStatement should contain arch %s", runtime.GOARCH)
	}
}

func TestGetVersionInfo(t *testing.T) {
	info := version.GetVersionInfo()
	if info == nil {
		t.Fatal("GetVersionInfo() returned nil")
	}
	if info.Major != version.Major {
		t.Errorf("Major mismatch: expected %d, got %d", version.Major, info.Major)
	}
	if info.Minor != version.Minor {
		t.Errorf("Minor mismatch: expected %d, got %d", version.Minor, info.Minor)
	}
	if info.Patch != version.Patch {
		t.Errorf("Patch mismatch: expected %d, got %d", version.Patch, info.Patch)
	}
	if info.OS != runtime.GOOS {
		t.Errorf("OS mismatch: expected %s, got %s", runtime.GOOS, info.OS)
	}
	if info.Arch != runtime.GOARCH {
		t.Errorf("Arch mismatch: expected %s, got %s", runtime.GOARCH, info.Arch)
	}
	// GoVersion is populated via runtime.Version() which always starts with "go"
	if !strings.HasPrefix(info.GoVersion, "go") {
		t.Errorf("GoVersion should start with 'go', got %s", info.GoVersion)
	}
}

func TestVersionInfoString(t *testing.T) {
	info := version.GetVersionInfo()
	str := info.String()
	if str == "" {
		t.Error("VersionInfo.String() should not return empty string")
	}
	if !strings.Contains(str, fmt.Sprintf("%d", version.Major)) {
		t.Errorf("String() should contain major version, got %s", str)
	}
	// Also verify the minor version is present in the string representation
	if !strings.Contains(str, fmt.Sprintf("%d", version.Minor)) {
		t.Errorf("String() should contain minor version, got %s", str)
	}
	// Patch version should also appear in the string representation
	if !strings.Contains(str, fmt.Sprintf("%d", version.Patch)) {
		t.Errorf("String() should contain patch version, got %s", str)
	}
}

// TestVersionIsNotEmpty is a quick sanity check to ensure the version string
// is non-empty and follows the expected semver-like format (X.Y.Z).
func TestVersionIsNotEmpty(t *testing.T) {
	if version.Version == "" {
		t.Error("Version string should not be empty")
	}
	parts := strings.Split(version.Version, ".")
	if len(parts) != 3 {
		t.Errorf("Version string should have 3 parts (major.minor.patch), got %q", version.Version)
	}
}
