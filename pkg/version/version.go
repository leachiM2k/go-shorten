package version

import "fmt"

var (
	// Version of the current build
	Version = "n/a"
	// GitCommit of the current build
	GitCommit = "dirty"
)

// GetInfo provides a ready to use version string
func GetInfo() string {
	return fmt.Sprintf("version: %s, commit: %s", Version, GitCommit)
}
