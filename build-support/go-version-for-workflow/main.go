package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	versionStr := strings.TrimPrefix(runtime.Version(), "go")
	if len(strings.Split(versionStr, ".")) == 2 {
		versionStr += ".0"
	}

	fmt.Println(versionStr)
}
