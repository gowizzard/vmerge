package version_core

import "regexp"

// core is to save the major number from the version core
var core = Core{}

// Split is to check the given version as a string
// The function checks the string with regex
// And split the value in major, minor and patch
func Split(version string) (Core, error) {

	regex := regexp.MustCompile(`[*]|\d+`)

	find := regex.FindAllString(version, 1)
	for index, value := range find {
		switch index {
		case 0:
			core.Major = value
		}
	}

	return core, nil

}
