package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Users_struc1 struct {
	ID     int    `json: "id"`
	Name   string `json: "name"`
	Email  string `json: "email"`
	Gender string `json: "gender"`
	Status string `json: "status"`
}

type Users_struc2 struct {
	ID          int    `json: "id"`
	Name        string `json: "name"`
	Email       string `json: "email"`
	Gender      string `json: "gender"`
	Status      string `json: "status"`
	Post_amount int    `json: "post_amount"`
	Post        []Post `json: "post"`
}
type Post struct {
	ID             int       `json: "id"`
	User_id        int       `json: "id"`
	Title          string    `json: "title"`
	Body           string    `json: "body"`
	Comment_amount int       `json: "comment_amount"`
	Comment        []Comment `json: "comment"`
}
type Comment struct {
	ID      int    `json: "id"`
	Post_id int    `json: "post_id"`
	Name    string `json: "name"`
	Email   string `json: "email"`
	Body    string `json: "body"`
}

func sentApiGetData(url string) (string, error) {
	// Create an HTTP client
	client := http.DefaultClient

	// Send a GET request to the URL
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the response body to a string
	responseString := string(body)
	return responseString, nil
}

var users_1 []Users_struc1
var users_2 []Users_struc2
var posts []Post
var comments []Comment

func main() {
	url := "https://gorest.co.in/public/v2/users"
	url_post := "https://gorest.co.in/public/v2/posts"
	url_comment := "https://gorest.co.in/public/v2/comments"

	responseString, err := sentApiGetData(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err_0 := json.Unmarshal([]byte(responseString), &users_1)
	if err_0 != nil {
		fmt.Println("Error:", err_0)
		return
	}

	responseString2, err := sentApiGetData(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err_1 := json.Unmarshal([]byte(responseString2), &users_2)
	if err_1 != nil {
		fmt.Println("Error:", err_1)
		return
	}

	responseString3, err := sentApiGetData(url_post)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err_2 := json.Unmarshal([]byte(responseString3), &posts)
	if err_2 != nil {
		fmt.Println("Error:", err_2)
		return
	}

	responseString4, err := sentApiGetData(url_comment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err_3 := json.Unmarshal([]byte(responseString4), &comments)
	if err_3 != nil {
		fmt.Println("Error:", err_3)
		return
	}

	for index_posts := range posts {
		post_pointer := &posts[index_posts]
			for index_comment := range comments {
				comment := comments[index_comment]
				if post_pointer.ID == comment.Post_id {
					post_pointer.Comment = append(post_pointer.Comment, comment)
					post_pointer.Comment_amount += 1
					fmt.Println("In user comment case", post_pointer.ID)
					fmt.Println("index_posts :", index_posts)
					fmt.Println("index_comment :", index_comment)
					fmt.Println("##########################")

				}
			}
		}

	for i := range users_2 {
		user := &users_2[i]

		for j := range posts {
			post := posts[j]
			if user.ID == post.User_id {
				user.Post = append(user.Post, post)
				user.Post_amount += 1
				fmt.Println("In user post case", post.User_id)
				fmt.Println("I :", i)
				fmt.Println("J :", j)
				fmt.Println("##########################")

			}
		}
	}

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/api/get-user-struct1", getUserStruct1).Methods("GET")
	router.HandleFunc("/api/get-user-struct2", getUserStruct2).Methods("GET")
	log.Println("Server listening on port 8080")
	trr := http.ListenAndServe(":8080", router)
	if trr != nil {
		log.Fatal(trr)
	}

	fmt.Println("Hello world")
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}
func getUserStruct1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Filter the users with status "active"
	activeUsers := make([]Users_struc1, 0)
	for _, user := range users_1 {
		if user.Status == "active" {
			activeUsers = append(activeUsers, user)
		}
	}
	json.NewEncoder(w).Encode(activeUsers)
}
func getUserStruct2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Filter the users with status "active"
	json.NewEncoder(w).Encode(users_2)
}
