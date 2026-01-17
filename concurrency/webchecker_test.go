package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebstiteChecker(url string) bool {
	return url != "http://bye.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://hi.com",
		"http://hello.com",
		"http://bye.com",
	}

	got := CheckWebsites(mockWebstiteChecker, websites)

	want := map[string]bool{
		"http://hi.com":    true,
		"http://hello.com": true,
		"http://bye.com":   false,
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowMockWebstiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return url != "http://bye.com"
}

func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 100)
	for i := range 100 {
		urls[i] = "http://hi.com"
	}

	for b.Loop() {
		CheckWebsites(slowMockWebstiteChecker, urls)
	}
}
