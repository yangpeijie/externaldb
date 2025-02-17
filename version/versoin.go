package version

// Version Version
const (
	version = "1.4.3"
)

// GitCommit git Version
var GitCommit string

// ReleaseDate ...
var ReleaseDate string

//GetVersion 获取版本信息
func GetVersion() string {
	if GitCommit != "" {
		return version + "-" + ReleaseDate + "-" + GitCommit
	}
	return version
}
