package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://wtf.com.com.com" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"https://www.twitter.com",
		"https://wtf.com.com.com",
	}

	want := map[string]bool{
		"https://www.google.com":  true,
		"https://www.twitter.com": true,
		"https://wtf.com.com.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}
