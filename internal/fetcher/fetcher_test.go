package fetcher

import (
	"testing"
)

func TestFetch(t *testing.T) {
	if err := Fetch(); err != nil {
		t.Error(err)
	}
}
