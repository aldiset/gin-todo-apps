package main

import (
   "io/ioutil"
   "log"
   "net/http"
)

func main() {

	for i := 0;i<1001;i++ {
		resp, err := http.Get("http://localhost:3030/todo-items?activity_group_id=3")
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		log.Printf(sb)
	}	
}