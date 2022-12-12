package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Posts    []Post
	Comments []Comment
	Profile  Profile
}

type Post struct {
	Id    int
	Title string
}

type Comment struct {
	Id     int
	Body   string
	PostId int
}

type Profile struct {
	Name string
}

func main() {
	println(GetData())
	println(PostData())
	println(PutData())
	println(DeleteData())
}

// Create Operation : Post Request
func PostData() string {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "amar",
		"email": "amar123@abc.com",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	if err != nil {
		println(err)
	}
	return resp.Status
}

// Read Operation : Get Request
func GetData() string {
	resp, err := http.Get("https://my-json-server.typicode.com/typicode/demo/db")
	if err != nil {
		print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		print(err)
	}
	fmt.Println(resp.Status)
	//print whole object
	print(string(body))

	// print single object
	var response Response
	json.Unmarshal(body, &response)

	// print Posts Object
	for _, p := range response.Posts {
		println("id:", p.Id, ", title: ", p.Title)
	}
	println("=======")

	// Print Comments Object
	for _, c := range response.Comments {
		println("id:", c.Id, ", body: ", c.Body, ", postId: ", c.PostId)
	}
	println("=======")

	//Print Profile Object
	println("name:", response.Profile.Name)

	return string(resp.Status)
}

// Update Operation : Put Request
func PutData() string {
	putbody, err := json.Marshal(map[string]string{
		"name":  "1mar",
		"email": "amar123@abcd.com",
	})
	if err != nil {
		println(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "https://postman-echo.com/put", bytes.NewBuffer(putbody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	return string(resp.Status)
}

// Delete Operation : Delete Request
func DeleteData() string {
	deleteBody, err := json.Marshal(map[string]string{
		"name":  "amar",
		"email": "amar123@abcd.com",
	})
	if err != nil {
		println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, "https://postman-echo.com/delete", bytes.NewBuffer(deleteBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	return string(resp.Status)
}
