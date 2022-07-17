package practice

import (
	"context"

	"log"
	"strings"
	"sync"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"

)

// golang 학습 테스트
func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {

		str := "Ann,Jenny,Tom,Zico"
		// TODO str을 , 단위로 잘라주세요.
		actual := strings.Split(str, ",")
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}
		//TODO assert 문을 활용해 actual과 expected를 비교해주세요.
		assert.Equal(t, expected, actual)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		var numbers [100]int
		var wg sync.WaitGroup
		wg.Add(100)
		for i := 0; i < 100; i++ {
			// 익명함수로 파라미터 i를 전달해야 모든 숫자가 추가될수 있다.
			go func(i int) {
				// TODO numbers에 i 값을 추가해보세요.
				//numbers = append(numbers, i) <- 완전히 랜덤값이 중복으로 들어오므로 fail
				numbers[i] = i

				//Done이 될때마다 Add 카운트는 -1
				wg.Done()
			}(i)
		}
		// 0이 될때까지 기다림
		wg.Wait()

		// actual : [0 1 2 ... 99]
		var expected []int = make([]int, 100)
		// TODO expected를 만들어주세요.
		for i := 0; i < 100; i++ {
			expected[i] = i
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
		outputCh := make(chan int, 3)
		go func(inputCh <-chan int) {
			for {
				select {
				case value, ok := <-inputCh:

					outputCh <- value * 10
					if !ok {
						//log.Panic("input channel closed") Panic을 쓰면 아래 테스트케이스에서 Fail처리됨...
						log.Println("input channel closed")
						close(outputCh)
						return
					}
				}
			}
		}(inputCh)

		var actual []int
		for i := 0; i < 3; i++ {
			actual = append(actual, <-outputCh)
		}

		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		ctx := context.TODO()
		// TODO 3초후에 종료하는 timeout context로 만들어주세요.
		ctx, cancel := context.WithTimeout(ctx, add)
		//두 번째 파라미터로 전달한 duration이 지나면  여기선 cancle(), 컨텍스트에 취소 신호가 전달된다.
		defer cancel()


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
		ctx := context.TODO()
		// TODO 3초후에 종료하는 timeout context로 만들어주세요.
		// 두 번째 파라미터로 time.Time 값을 받는데, 이 시간이 되면 컨텍스트에 취소 신호가 전달된다
		ctx, cancel := context.WithDeadline(ctx, startTime.Add(add))
		defer cancel()

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		// context에 key, value를 추가해보세요.
		ctx := context.TODO()
		ctx = context.WithValue(ctx, "job", "devloper")
		// 추가된 key, value를 호출하여 assert로 값을 검증해보세요.
		assert.Equal(t, ctx.Value("job"), "devloper")
		// 추가되지 않은 key에 대한 value를 assert로 검증해보세요.
		assert.Nil(t, ctx.Value("name"))

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
