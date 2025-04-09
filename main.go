package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type McpExample struct{}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Greets the individal with the provided name
func (m *McpExample) GreetSomeone(name string) string {
	return "Hello, " + name + "!"
}

// Get all of the names of the posts as a JSON string
func (m *McpExample) GetAllPosts() (string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	posts := []Post{}
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return "", err
	}

	b, err := json.Marshal(posts)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Get a single post by ID
// If the post is not found, return an error
func (m *McpExample) GetPostByID(id int) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("post with ID %d not found", id)
	}

	post := Post{}
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return "", err
	}

	b, err := json.Marshal(post)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
