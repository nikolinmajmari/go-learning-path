package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "xapi://xapi.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://a.com",
		"xapi://xapi.com",
		"ws://a.com",
	}
	want := map[string]bool{
		"https://a.com":   true,
		"xapi://xapi.com": false,
		"ws://a.com":      true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v got %v", want, got)
	}

}

func slowSetupWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowSetupWebsiteChecker, urls)
	}
}
