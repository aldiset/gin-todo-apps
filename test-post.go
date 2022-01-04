package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for i:=0;i<100;i++{
		httpposturl := "http://localhost:3030/todo-items"
		fmt.Println("HTTP JSON POST URL:", httpposturl)

		var jsonData = []byte(`{
			"activity_group_id": 3,
			"title": "item 5.4"
			}`)
		request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()

		fmt.Println("response Status:", response.Status)
		fmt.Println("response Headers:", response.Header)
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("response Body:", string(body))

	}
}
