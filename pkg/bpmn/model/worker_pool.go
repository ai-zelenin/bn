package model

import (
	"context"
	"sync"
)

type WorkerPayload[T any] func(ctx context.Context, wp *WorkerPool[T], val T)
type WorkerPool[T any] struct {
	workers int
	payload WorkerPayload[T]
	input   chan T
}

func NewWorkerPool[T any](workers int, payload WorkerPayload[T]) *WorkerPool[T] {
	return &WorkerPool[T]{
		workers: workers,
		payload: payload,
		input:   make(chan T, workers*10),
	}
}

func (w *WorkerPool[T]) Start(ctx context.Context) {
	wg := new(sync.WaitGroup)
	for i := 0; i < w.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-w.input:
					if !ok {
						return
					}
					w.payload(ctx, w, val)
				}
			}
		}()
	}
	wg.Wait()
	return
}

func (w *WorkerPool[T]) Handle(event T) {
	w.input <- event
}
