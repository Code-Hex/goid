package goid_test

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
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
			if want, got := goIDByStack(), goid.Get(); want != got {
				panic(fmt.Sprintf("want %d, but got %d", want, got))
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
}

func goIDByStack() int64 {
	buf := make([]byte, 64)
	s := buf[:runtime.Stack(buf, false)]
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	i, _ := strconv.ParseInt(string(s), 10, 64)
	return i
}
