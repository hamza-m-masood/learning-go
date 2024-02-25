package client

import (
	"fmt"
	"io"
	"net/http"
)
func Get(){
	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil{
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(body))


}
