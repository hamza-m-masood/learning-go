package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func Post(){
	client := &http.Client{}
	type PostBody struct {
		Title string `json:"title"`
		Body string `json:"Body"`
		UserId int `json:"userID"`
	}
	body := PostBody{
		Title: "foo",
    Body: "bar",
    UserId: 1,
	}
	bodyBytes, err := json.Marshal(&body)
	if err != nil{
		fmt.Println("error from marshalling: ", err)
	}
	reader := bytes.NewReader(bodyBytes)
	resp, err := client.Post("https://jsonplaceholder.typicode.com/posts", "application/json", reader)
	if err != nil{
		fmt.Println("error from marshalling: ", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("error from marshalling: ", err)
	}
	fmt.Println(string(respBody))
}
