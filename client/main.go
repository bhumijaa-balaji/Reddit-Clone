// client.go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const baseURL = "http://127.0.0.1:8223"

type Post struct {
    ID        string   `json:"id"`
    Title     string   `json:"title"`
    Content   string   `json:"content"`
    Author    string   `json:"author"`
    Subreddit string   `json:"subreddit"`
    Upvotes   int      `json:"upvotes"`
    Comments  []string `json:"comments"`
}

func main() {
    // Create a post
    newPost := Post{
        ID:        "1",
        Title:     "First Post",
        Content:   "This is the content of the first post",
        Author:    "user1",
        Subreddit: "golang",
        Upvotes:   0,
        Comments:  []string{},
    }
    createPost(newPost)

    // Get all posts
    getPosts()

    // Get a specific post
    getPost("1")

    // Update a post
    updatedPost := Post{
        ID:        "1",
        Title:     "Updated First Post",
        Content:   "This is the updated content of the first post",
        Author:    "user1",
        Subreddit: "golang",
        Upvotes:   5,
        Comments:  []string{"Great post!"},
    }
    updatePost("1", updatedPost)

    // Delete a post
    deletePost("1")
}

func createPost(post Post) {
    postJSON, _ := json.Marshal(post)
    resp, err := http.Post(baseURL+"/posts", "application/json", bytes.NewBuffer(postJSON))
    if err != nil {
        fmt.Println("Error creating post:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Created post:", string(body))
}

func getPosts() {
    resp, err := http.Get(baseURL + "/posts")
    if err != nil {
        fmt.Println("Error getting posts:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("All posts:", string(body))
}

func getPost(id string) {
    resp, err := http.Get(baseURL + "/posts/" + id)
    if err != nil {
        fmt.Println("Error getting post:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Post details:", string(body))
}

func updatePost(id string, post Post) {
    postJSON, _ := json.Marshal(post)
    req, _ := http.NewRequest(http.MethodPut, baseURL+"/posts/"+id, bytes.NewBuffer(postJSON))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error updating post:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Updated post:", string(body))
}

func deletePost(id string) {
    req, _ := http.NewRequest(http.MethodDelete, baseURL+"/posts/"+id, nil)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error deleting post:", err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Delete response:", string(body))
}
