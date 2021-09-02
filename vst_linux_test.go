// +build !plugin

package vst2_test

import (
	"fmt"
	"testing"

	"github.com/astonm/vst2"
)

func TestLinux(t *testing.T) {
	fmt.Printf("linux paths: %v\n", vst2.ScanPaths())
}
