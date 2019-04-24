package utils

import (
	"fmt"
	"os"
)

func Assert(cond bool, errMsg string) {
	if !cond {
		fmt.Printf("error: %s\n", errMsg)
		os.Exit(1)
	}
}
