package goConcurrency

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type MyContextKey string

func ContextValue() {
	ctx := context.WithValue(context.Background(), "title", "Go")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		if v := c.Value("title"); v != nil {
			fmt.Println("Found value:", v)
			return
		}
		fmt.Println("Key Not found", "title")
	}(ctx)
	wg.Wait()
}
func ContextValueDeep() {
	ctxOne := context.WithValue(context.Background(), MyContextKey("title"), "Go One")
	ctxTwo := context.WithValue(ctxOne, "key", "Go Two")
	ctxThree := context.WithValue(ctxTwo, "key", "Go Three")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		if v := c.Value(MyContextKey("title")); v != nil {
			fmt.Println("Found value:", v)
			return
		}
		fmt.Println("Key Not found", "title")
	}(ctxThree)
	wg.Wait()
}
func ContextCancelDeep() {
	//ctxOne, _ := context.WithCancel(context.Background())
	//ctxTwo, cancel := context.WithCancel(ctxOne)
	//ctxThree, _ := context.WithCancel(ctxOne)
	//ctxFour, _ := context.WithCancel(ctxTwo)
	ctxOne, _ := context.WithTimeout(context.Background(), time.Second*1)
	ctxTwo, cancel := context.WithTimeout(ctxOne, time.Second*1)
	ctxThree, _ := context.WithTimeout(ctxOne, time.Second*1)
	ctxFour, _ := context.WithTimeout(ctxTwo, time.Second*1)
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func(c context.Context) {
		defer wg.Done()
		select {
		case <-c.Done():
			fmt.Println("one canceled")
		}
	}(ctxOne)
	go func(c context.Context) {
		defer wg.Done()
		select {
		case <-c.Done():
			fmt.Println("two canceled")
		}
	}(ctxTwo)
	go func(c context.Context) {
		defer wg.Done()
		select {
		case <-c.Done():
			fmt.Println("three canceled")
		}
	}(ctxThree)
	go func(c context.Context) {
		defer wg.Done()
		select {
		case <-c.Done():
			fmt.Println("four canceled")
		}
	}(ctxFour)
	cancel()
	wg.Wait()

}
func ContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 4; i++ {
		go func(c context.Context, n int) {
			for {
				select {
				case <-c.Done():
					return
				default:
				}
				fmt.Println(strings.Repeat("  ", n), n)
				time.Sleep(time.Millisecond * 300)
			}
		}(ctx, i)
	}
	select {
	case <-time.NewTimer(time.Second * 3).C:
		cancel()
	}
	select {
	case <-ctx.Done():
		fmt.Println("context cancel")
	}

}
func ContextCancelTime() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	for i := 0; i < 4; i++ {
		go func(c context.Context, n int) {
			for {
				select {
				case <-c.Done():
					return
				default:
				}
				fmt.Println(strings.Repeat("  ", n), n)
				time.Sleep(time.Millisecond * 300)
			}
		}(ctx, i)
	}

	select {
	case <-time.NewTimer(time.Second * 1).C:
		cancel()
		fmt.Println("call cancel() Cancel")
	case <-ctx.Done():
		fmt.Println("main cancel")
	}
}
