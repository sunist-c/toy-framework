package test

import (
	"sync"
	"testing"

	"github.com/sunist-c/toy-framework/ds/channel"
)

func TestChannel(t *testing.T) {
	t.Run("channel test", func(t *testing.T) {
		ch := channel.NewChannel[int]()
		for i := 0; i < 10086; i++ {
			ch.Push(i)
		}

		for i := 0; i < 10090; i++ {
			x, s := ch.Pop()
			if i < 10086 && x != i {
				t.Errorf("expected %d, got %d", i, x)
			}
			if i < 10086 && !s {
				t.Errorf("expected failed")
			}

			if i >= 10086 && s {
				t.Errorf("expected unfailed")
			}
		}
	})
}

func BenchmarkTestChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := channel.NewChannel[int]()
		wWg := &sync.WaitGroup{}
		rWg := &sync.WaitGroup{}
		wWg.Add(100 * 10000)
		rWg.Add(100 * 10000)
		for j := 0; j < 100; j++ {
			go func() {
				for k := 0; k < 10000; k++ {
					ch.Push(k)
					wWg.Done()
				}
			}()
		}

		wWg.Wait()

		for j := 0; j < 100; j++ {
			for k := 0; k < 10000; k++ {
				_, s := ch.Pop()
				if !s {
					b.Errorf("error")
				}
				rWg.Done()
			}
		}
		rWg.Wait()
	}
}
