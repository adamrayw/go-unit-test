package helper

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Windows")
	}

	fmt.Println("Test Running")
}
