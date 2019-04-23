package utils

import (
	"fmt"
	"os"
)

func Assert(cond bool, errMsg string) {
	if !cond {
		fmt.Errorf("error: %s", errMsg)
		os.Exit(1)
	}
}
