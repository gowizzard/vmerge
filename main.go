package main

import (
	"fmt"
	"github.com/gowizzard/vmerge/version_core"
	"os"
)

// version, suffix, output are to save the environment and output data
var (
	version = os.Getenv("VERSION")
	suffix  = os.Getenv("SUFFIX")
	output  []byte
)

// main is there to start the solution
// And return the output to command line
func main() {

	core, err := version_core.Split(version)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	output = fmt.Append(output, "::set-output")
	output = fmt.Appendf(output, "%-*s", 1, "")
	output = fmt.Appendf(output, "name=branch_name::%s%s", suffix, core.Major)

	fmt.Printf("%s", output)

}
