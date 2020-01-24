package architecture

import (
	"runtime"
	"testing"
)

func TestArchitecture(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Does not work in amd64 architecture")
	}

	// ...
}
