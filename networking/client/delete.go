package client

import (
	"fmt"
	"io"
	"net/http"
)

func Delete(){
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/delete", nil)
	if err!= nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err!= nil{
		fmt.Println(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(string(respBody))
}
