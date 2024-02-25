package client

import (
	"fmt"
	"io"
	"net/http"
)
func Get(){
	client := &http.Client{}

	resp, err := client.Get("http://localhost:8080")
	if err != nil{
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(body))


}
