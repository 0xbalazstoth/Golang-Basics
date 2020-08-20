package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bdy struct {
	title  string
	body   string
	userId string
	id     int
}

func main() {
	req, err := json.Marshal(map[string]string{
		"title":  "Title100",
		"body":   "Body100",
		"userId": "6565",
	})

	if err != nil {
		log.Fatalln(err)
	}

	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(req))
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	var jsonBdy Bdy
	json.Unmarshal([]byte(string(body)), &jsonBdy)
	fmt.Printf("Title: %v\nBody: %v\nuserId: %v\n", jsonBdy.title, jsonBdy.body, jsonBdy.userId)
}
