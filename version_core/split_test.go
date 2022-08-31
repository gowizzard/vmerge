package version_core_test

import (
	"github.com/gowizzard/vmerge/version_core"
	"log"
	"reflect"
	"testing"
)

// TestSplit is to check the version core split function
// We simulate the version input and check the result
func TestSplit(t *testing.T) {

	tests := []struct {
		version  string
		expected version_core.Core
	}{
		{
			version: "v1.2.4",
			expected: version_core.Core{
				Major: "1",
			},
		},
		{
			version: "v3.7.0-alpha.2+testing-12345a",
			expected: version_core.Core{
				Major: "3",
			},
		},
		{
			version: "0.12.0",
			expected: version_core.Core{
				Major: "0",
			},
		},
	}

	for _, value := range tests {

		core, err := version_core.Split(value.version)
		if err != nil {
			log.Fatalln(err)
		}

		numbers := []struct {
			Number1 string
			Number2 string
		}{
			{
				Number1: value.expected.Major,
				Number2: core.Major,
			},
		}

		for _, value := range numbers {
			if !reflect.DeepEqual(value.Number1, value.Number2) {
				t.Fatalf("expected: %s, got %s", value.Number1, value.Number2)
			}
		}

	}

}
