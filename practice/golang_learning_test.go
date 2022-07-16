package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

// golang 학습 테스트
func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		str := "Ann,Jenny,Tom,Zico"
		actual := strings.Split(str, ",")
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}
		assert.EqualValues(t, expected, actual)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		numbers := make([]int, 100)
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			i2 := i
			go func() {
				defer wg.Done()
				numbers[i2] = i2
			}()
		}

		var expected []int // actual : [0 1 2 ... 99]
		for j := 0; j < 100; j++ {
			expected = append(expected, j)
		}
		wg.Wait()
		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {

		inputCh := generate()
		outputCh := make(chan int)

		go func() {
			for {
				select {
				case value, ok := <-inputCh:
					if !ok {
						close(outputCh)
						return
					}
					outputCh <- value * 10
				}
			}
		}()

		var actual []int
		for value := range outputCh {
			actual = append(actual, value)
		}
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx, _ := context.WithTimeout(context.Background(), add)

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context deadline", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx, _ := context.WithDeadline(context.Background(), startTime.Add(add))

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "key", "value")
		assert.Equal(t, "value", ctx.Value("key"))
		assert.Equal(t, "august", ctx.Value("key"))
	})
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	return ch
}
