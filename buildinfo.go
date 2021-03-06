package buildinfo

import (
	"bytes"
	"fmt"
	"runtime"
)

//noinspection GoUnusedConst
const (
	// ShortFlag is the default flag to print the current build information of the app.
	ShortFlag = "v"
	// LongFlag is an alternative long version that may be used together with ShortFlag.
	LongFlag = "version"

	// MetricName is a default name for the metric (without namespace).
	MetricName = "build_info"
	// MetricHelp is a default help text that describes the metric.
	MetricHelp = "Metric with build information labels and a constant value of '1'."
)

// BuildInfo
type BuildInfo struct {
	Version  string
	Revision string
	Branch   string
	Date     string
}

// GoVersion returns the version of the used Go runtime. See `runtime.Version()`
// for additional details.
func (bld BuildInfo) GoVersion() string { return runtime.Version() }

// String returns the string representation of the build information.
// It includes release version, revision and build date.
func (bld BuildInfo) String() string {
	return fmt.Sprintf("%s, #%s @ %s", bld.Version, bld.Revision, bld.Date)
}

// Map returns the build information as a map. Field names are lowercase.
// Empty fields within BuildInfo are omitted.
func (bld BuildInfo) Map() map[string]string {
	m := make(map[string]string, 5)
	m["version"] = bld.Version

	if bld.Revision != "" {
		m["rev"] = bld.Revision
	}
	if bld.Branch != "" {
		m["branch"] = bld.Branch
	}
	if bld.Date != "" {
		m["date"] = bld.Date
	}

	m["goversion"] = runtime.Version()
	return m
}

// MarshalJSON returns valid JSON output.
// Empty fields within BuildInfo are omitted.
func (bld BuildInfo) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(`{"version":"`)
	buf.WriteString(bld.Version)

	if bld.Revision != "" {
		buf.WriteString(`","rev":"`)
		buf.WriteString(bld.Revision)
	}
	if bld.Branch != "" {
		buf.WriteString(`","branch":"`)
		buf.WriteString(bld.Branch)
	}
	if bld.Date != "" {
		buf.WriteString(`","date":"`)
		buf.WriteString(bld.Date)
	}

	buf.WriteString(`","goversion":"`)
	buf.WriteString(runtime.Version())
	buf.WriteString(`"}`)

	return buf.Bytes(), nil
}

const (
	DummyVersion  = "0.0.0"
	DummyRevision = "abcdef"
	DummyBranch   = "HEAD"
	DummyDate     = "1997-08-29 13:37:00"
)

// IsDummy returns `true` when all fields' values within a `BuildInfo` are dummy
// values. This may indicate the build information variables are not properly
// overwritten when a new build is made.
func IsDummy(bld BuildInfo) bool {
	return bld.Version == DummyVersion &&
		bld.Revision == DummyRevision &&
		bld.Branch == DummyBranch &&
		bld.Date == DummyDate
}
