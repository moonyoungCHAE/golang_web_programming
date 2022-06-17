package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestGolang(t *testing.T) {

	t.Run("string test", func(t *testing.T) {
		str := "Ann,Jenny,Tom,Zico"
		actual := strings.Split(str, ",")
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}
		assert.ElementsMatch(t, actual, expected)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		var numbers []int
		go func() {
			for i := 0; i < 100; i++ {
				numbers = append(numbers, i)
			}
		}()

		var expected []int
		for i := 0; i < 100; i++ {
			expected = append(expected, i)
		}

		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {
		/*
			TODO 주어진 코드를 수정해서 테스트가 통과하도록 해주세요!

			- inputCh에 1, 2, 3 값을 넣는다.
			- inputCh로 부터 값을 받아와, value * 10 을 한 후 outputCh에 값을 넣어준다.
			- outputCh에서 읽어온 값을 비교한다.
		*/

		inputCh := generate()
		outputCh := make(chan int)
		go func() {
			defer close(outputCh)
			for {
				select {
				case value, ok := <-inputCh:
					if !ok {
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
		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, time.Second*3)

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
		ctx := context.Background()
		ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second*3))

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "my_key", "my_value")
		assert.Equal(t, "my_value", ctx.Value("my_key"))
		assert.Nil(t, ctx.Value("your key"))
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
