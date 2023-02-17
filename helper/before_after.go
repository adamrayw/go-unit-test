package helper

import (
	"fmt"
	"testing"
)

func TestBeforeAfter(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}
