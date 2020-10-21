package goid_test

import (
	"bytes"
	"runtime"
	"sync"
	"testing"

	"github.com/Code-Hex/goid"
)

func TestA(t *testing.T) {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	n := 3
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			t.Log(goid.Get())
			checkGoID(t)
			t.Log("--------------")
			mu.Unlock()
		}()
	}
	wg.Wait()
}

func checkGoID(t *testing.T) {
	buf := make([]byte, 64)
	s := buf[:runtime.Stack(buf, false)]
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	t.Log(string(s))
}
