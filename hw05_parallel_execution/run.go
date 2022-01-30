package hw05parallelexecution

import (
	"errors"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/dlsniper/debugger"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}

	activeTasksChannel := make(chan Task, n)
	maxErrorsCount := int32(m)
	var errorsCount int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			debugger.SetLabels(func() []string {
				return []string{
					"worker", strconv.Itoa(i),
					"__", "__",
				}
			})
			defer wg.Done()
			for task := range activeTasksChannel {
				if atomic.LoadInt32(&errorsCount) >= maxErrorsCount {
					break
				}

				if err := task(); err != nil {
					atomic.AddInt32(&errorsCount, 1)
				}
			}
		}(i)
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errorsCount) >= maxErrorsCount {
			break
		}
		activeTasksChannel <- task
	}
	close(activeTasksChannel)
	wg.Wait()

	if errorsCount >= maxErrorsCount {
		return ErrErrorsLimitExceeded
	}

	return nil
}
