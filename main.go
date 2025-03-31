package main

import (
	"io"
	"net/http"
)

type McpExample struct{}

// Greets the individal with the provided name
func (m *McpExample) GreetSomeone(name string) string {
	return "Hello, " + name + "!"
}

func (m *McpExample) GetAllPosts() (string, error) {
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
