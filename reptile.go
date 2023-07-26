package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Comment struct {
	Email string `json:"email"`
}

func main() {
	var wg sync.WaitGroup
	emails := make(chan string, 500)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/comments", i)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			var comments []Comment
			json.Unmarshal(body, &comments)
			for _, comment := range comments {
				emails <- comment.Email
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(emails)
	}()

	file, err := os.Create("emails.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for email := range emails {
		fmt.Println(email)
		file.WriteString(email + "\n")
	}
}
