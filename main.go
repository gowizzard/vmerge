package main

import (
	"fmt"
	"github.com/gowizzard/vmerge/version_core"
	"os"
)

var (
	version = os.Getenv("VERSION")
	suffix  = os.Getenv("SUFFIX")
	output  []byte
)

func main() {

	core, err := version_core.Split(version)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	output = fmt.Append(output, "::set-output")
	output = fmt.Appendf(output, "%s", " ")
	output = fmt.Appendf(output, "name=branch::%s%s", suffix, core)

	fmt.Print(string(output))

}
