package update

// Version is ...
type Version struct {
	CurrentVersion string `json:"current_version"`
	GitCommit      string `json:"gitCommit"`
	BuildDate      string `json:"buildDate"`
	NewVersion     string `json:"new,omitempty"`
	ReleaseURL     string `json:"release_url,omitempty"`
}

var versionInfo *Version

func SetVersion(ver *Version) {
	versionInfo = ver
}

func VersionInfo() *Version {
	return versionInfo
}
