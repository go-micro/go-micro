package redis

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/micro/go-micro/v3/cache"
)

var (
	nodes = []string{"127.0.0.1:6379", "127.0.0.1:6380", "127.0.0.1:6381"}
)

func TestNewCache(t *testing.T) {
	for _, nod := range nodes {
		cache.Nodes(nod)
	}

	cache := NewCache(cache.Nodes())

	t.Log(cache.String())
}

func TestSet(t *testing.T) {
	for _, nod := range nodes {
		cache.Nodes(nod)
	}

	cache := NewCache(cache.Nodes())

	key := "test"
	val := 100
	if err := cache.Set(key, 100); err != nil {
		t.Fatal(err.Error())
	}

	rs, err := cache.Get(key)
	if err != nil {
		t.Fatal(err.Error())
	}

	v, ok := rs.(string)
	fmt.Println(v, ok)

	if ok {
		vi, err := strconv.Atoi(v)
		if err != nil {
			t.Fatal(err.Error())
		}

		if val == vi {
			t.Logf("set(%v) get(%v) ok", val, vi)
		} else {
			t.Fatalf("set(%v) get(%v) fail", val, vi)
		}
	}
}