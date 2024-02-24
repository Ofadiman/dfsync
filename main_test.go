package main

import (
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestMain(m *testing.M) {
	v := m.Run()

	// Cleanup obsolete snapshot tests after all tests were run.
	snaps.Clean(m, snaps.CleanOpts{Sort: true})

	os.Exit(v)
}
