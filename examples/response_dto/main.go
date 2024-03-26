package main

import (
	"fmt"
	"github.com/WEG-Technology/room"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	response, err := room.NewRequest("https://jsonplaceholder.typicode.com/posts/1").Send()

	if err != nil {
		fmt.Println("Error:", err)
	}

	post := response.DTO(&Post{}).(*Post)

	fmt.Println("Response:", post.Title)
}
