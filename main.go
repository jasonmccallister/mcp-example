package main

import (
	"encoding/json"
	"net/http"
)

type McpExample struct{}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Greets the individal with the provided name
func (m *McpExample) GreetSomeone(name string) string {
	return "Hello, " + name + "!"
}

// Get all of the names of the posts
func (m *McpExample) GetAllPosts() ([]string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	posts := []Post{}
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}

	titles := []string{}
	for _, post := range posts {
		titles = append(titles, post.Title)
	}

	return titles, nil
}

// Returns the post with the provided ID
func (m *McpExample) GetPostByID(id int) (*Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + string(id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	post := &Post{}
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, err
	}

	return post, nil
}
