package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Example struct {
	Example string `json:"example"`
}

type ExampleRepo struct {
	Client http.Client
	URL    string
}

func (e *ExampleRepo) HttpPostExample() error {
	obj, _ := json.Marshal(Example{
		Example: "Hello World",
	})

	res, err := e.Client.Post(e.URL, "application/json", bytes.NewBuffer(obj))
	defer res.Body.Close()

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("oh no")
	}

	return nil
}

func (e *ExampleRepo) HttpGetExample() (*Example, error) {
	res, err := e.Client.Get(e.URL)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("oh no")
	}

	var receiver Example
	json.NewDecoder(res.Body).Decode(&receiver)

	return &receiver, nil
}

func (e *ExampleRepo) HttpDoExample() (*Example, error) {
	req, err := http.NewRequest("GET", e.URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Something", "something")
	resp, err := e.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("oh no")
	}

	var receiver Example
	json.NewDecoder(resp.Body).Decode(&receiver)

	return &receiver, nil
}

func (e *ExampleRepo) GoChannelFunc() string {
	say := func(s string, c chan string) {
		c <- s
	}

	c := make(chan string)
	go say("hello", c)
	go say("world", c)
	x, y := <-c, <-c

	return fmt.Sprintf("%s, %s", x, y)
}

func (e *ExampleRepo) GoChannelCloseWithResults() string {
	say := func(c chan string) {
		arr := []string{"hello", "world"}
		for i := 0; i < len(arr); i++ {
			c <- arr[i]
		}
		close(c)
	}
	c := make(chan string)
	go say(c)

	res := make([]string, 0, 2)
	for i := range c {
		res = append(res, i)
	}

	return strings.Join(res, ", ")
}

func (e *ExampleRepo) GoWaitGroup() string {
	arr := []string{"Hello", "World"}
	var group sync.WaitGroup
	group.Add(len(arr))
	strs := make(chan string, len(arr))
	for _, item := range arr {
		go func(i string) {
			defer group.Done()
			strs <- i
		}(item)
	}
	group.Wait()
	close(strs)

	res := make([]string, 0, len(arr))
	for v := range strs {
		res = append(res, v)
	}

	return strings.Join(res, ", ")
}

func (e *ExampleRepo) GoSimpleContext() (*string, error) {
	say := func(ctx context.Context, ch chan string) {
		time.Sleep(time.Second * 1)
		ch <- "Hello World"
	}

	ch := make(chan string, 1)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go say(ctxTimeout, ch)

	select {
	case <-ctxTimeout.Done():
		return nil, ctxTimeout.Err()
	case result := <-ch:
		return &result, nil
	}
}

func (e *ExampleRepo) GoHttpContext() (*Example, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, err := http.NewRequestWithContext(ctxTimeout, http.MethodGet, e.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var receiver Example
	json.NewDecoder(resp.Body).Decode(&receiver)

	return &receiver, nil
}

func (e *ExampleRepo) GoContextInForLoop() string {
	say := func(s string, ctx context.Context, wg *sync.WaitGroup, c chan<- string) {
		defer wg.Done()
		select {
		case c <- s:
			fmt.Println(s, "is in channel")
		case <-ctx.Done():
			fmt.Println(s, "has just been canceled", ctx.Err())
		}
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	arr := []string{"hello", "world"}
	ch := make(chan string, len(arr))
	for _, v := range arr {
		wg.Add(1)
		go say(v, ctx, &wg, ch)
	}

	wg.Wait()
	close(ch)
	cancel()

	res := make([]string, 0, len(ch))
	for v := range ch {
		res = append(res, v)
	}

	r := strings.Join(res, ", ")
	return r
}

func main() {
	fmt.Println("Hello World")
}
