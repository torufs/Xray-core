// Package buildinfo exposes compile-time and runtime metadata about the
// current Xray-core build.
//
// Build metadata (date, commit hash, branch) can be injected at link time
// using -ldflags, for example:
//
//	-ldflags "-X github.com/xtls/xray-core/common/buildinfo.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)
//	          -X github.com/xtls/xray-core/common/buildinfo.CommitHash=$(git rev-parse --short HEAD)
//	          -X github.com/xtls/xray-core/common/buildinfo.Branch=$(git rev-parse --abbrev-ref HEAD)"
//
// When not set via ldflags, each field defaults to the string "unknown".
// Runtime fields (GoVersion, OS, Arch) are always populated from the
// standard library at startup.
//
// Note: for local builds, you can use the Makefile target `make build` which
// automatically injects these ldflags. See the project README for details.
//
// Personal note: I also find it useful to inject a custom "Builder" tag so
// I can distinguish my local builds from official releases at a glance:
//
//	-X github.com/xtls/xray-core/common/buildinfo.Builder=myusername
//
// Tip: you can also inject BuildDate using RFC3339 format for easier parsing
// in scripts: $(date -u +%Y-%m-%dT%H:%M:%SZ)
package buildinfo
