package pokecache

import (
	"testing"
	"time"
)

func Test_reap_loop(t *testing.T) {
	basetime := time.Millisecond * 6
	cache := NewCache(basetime)

	cache.Add("lol", "lol")
	if _, ok := cache.Get("lol"); !ok {
		t.Errorf("expected to find key: lol")
	}

	time.Sleep(basetime / 2)
	cache.Add("hmm", "hmm")

	time.Sleep(basetime / 2)

	if _, ok := cache.Get("hmm"); !ok {
		t.Errorf("expected to find key: hmm")
	}
	if _, ok := cache.Get("lol"); ok {
		t.Errorf("expected to not find key: lol")
	}
}
