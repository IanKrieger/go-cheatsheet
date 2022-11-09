package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HTTP_Calls_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	repo := ExampleRepo{
		URL:    server.URL,
		Client: http.Client{},
	}

	value := repo.HttpPostExample()
	if value != nil {
		t.Errorf("Expected nil, got %s", value.Error())
	}
}

func Test_HTTP_Calls_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"example":"Hello World"}`))
	}))
	defer server.Close()

	repo := ExampleRepo{
		URL:    server.URL,
		Client: http.Client{},
	}

	ex, err := repo.HttpGetExample()
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}

	if ex.Example != "Hello World" {
		t.Errorf("Expected Hello World, got %s", ex.Example)
	}
}

func Test_HTTP_Calls_Do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"example":"Hello World"}`))
	}))
	defer server.Close()

	repo := ExampleRepo{
		URL:    server.URL,
		Client: http.Client{},
	}

	ex, err := repo.HttpDoExample()
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}

	if ex.Example != "Hello World" {
		t.Errorf("Expected Hello World, got %s", ex.Example)
	}
}

func Test_Go_Channel(t *testing.T) {
	repo := ExampleRepo{}
	res := repo.GoChannelFunc()

	if res != "hello, world" && res != "world, hello" {
		t.Errorf("Expected Hello World, got %s", res)
	}
}

func Test_Go_Channel_Close_Results(t *testing.T) {
	repo := ExampleRepo{}
	res := repo.GoChannelCloseWithResults()

	if res != "hello, world" {
		t.Errorf("Expected hello, world: got %s", res)
	}
}

func Test_Go_Wait_Group(t *testing.T) {
	repo := ExampleRepo{}
	res := repo.GoWaitGroup()

	if res != "Hello, World" && res != "World, Hello" {
		t.Errorf("Expected Hello, World: got %s", res)
	}
}

func Test_Go_Simple_Context(t *testing.T) {
	repo := ExampleRepo{}
	res, err := repo.GoSimpleContext()

	if err != nil {
		t.Errorf("Expected Hello, World: got %s", err.Error())
	}

	if *res != "Hello World" {
		t.Errorf("Expected Hello, World: got %s", *res)
	}
}

func Test_HTTP_Calls_Context(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"example":"Hello World"}`))
	}))
	defer server.Close()

	repo := ExampleRepo{
		URL:    server.URL,
		Client: http.Client{},
	}

	ex, err := repo.GoHttpContext()
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}

	if ex.Example != "Hello World" {
		t.Errorf("Expected Hello World, got %s", ex.Example)
	}
}

func Test_Go_Context_For_Loop(t *testing.T) {
	repo := ExampleRepo{}
	res := repo.GoContextInForLoop()

	if res != "hello, world" && res != "world, hello" {
		t.Errorf("Expected hello, world: got %s", res)
	}
}
