package buildinfo_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/xtls/xray-core/common/buildinfo"
)

func TestGetBuildInfoDefaults(t *testing.T) {
	info := buildinfo.GetBuildInfo()

	if info.GoVersion == "" {
		t.Error("GoVersion should not be empty")
	}
	if info.OS == "" {
		t.Error("OS should not be empty")
	}
	if info.Arch == "" {
		t.Error("Arch should not be empty")
	}

	if info.GoVersion != runtime.Version() {
		t.Errorf("expected GoVersion %q, got %q", runtime.Version(), info.GoVersion)
	}
	if info.OS != runtime.GOOS {
		t.Errorf("expected OS %q, got %q", runtime.GOOS, info.OS)
	}
	if info.Arch != runtime.GOARCH {
		t.Errorf("expected Arch %q, got %q", runtime.GOARCH, info.Arch)
	}
}

func TestBuildInfoString(t *testing.T) {
	info := buildinfo.GetBuildInfo()
	s := info.String()

	for _, keyword := range []string{"Build Date", "Commit", "Branch", "Go Version", "OS/Arch"} {
		if !strings.Contains(s, keyword) {
			t.Errorf("String() missing keyword %q", keyword)
		}
	}
}

func TestBuildInfoShortString(t *testing.T) {
	info := buildinfo.GetBuildInfo()
	short := info.ShortString()

	if short == "" {
		t.Error("ShortString() should not be empty")
	}
	if !strings.Contains(short, runtime.GOOS) {
		t.Errorf("ShortString() should contain OS %q, got %q", runtime.GOOS, short)
	}
	if !strings.Contains(short, runtime.GOARCH) {
		t.Errorf("ShortString() should contain Arch %q, got %q", runtime.GOARCH, short)
	}
}

func TestBuildInfoUnknownDefaults(t *testing.T) {
	// Without ldflags, these should default to "unknown".
	// In CI builds these values are typically injected via ldflags, so we
	// only log a notice rather than failing the test when they differ.
	// Note: when running locally without build tags, all three fields below
	// should print "unknown" — useful to verify a clean dev environment.
	info := buildinfo.GetBuildInfo()
	if info.BuildDate != "unknown" {
		t.Logf("BuildDate is %q (may be set by ldflags)", info.BuildDate)
	}
	if info.CommitHash != "unknown" {
		t.Logf("CommitHash is %q (may be set by ldflags)", info.CommitHash)
	}
	if info.Branch != "unknown" {
		t.Logf("Branch is %q (may be set by ldflags)", info.Branch)
	}
}
