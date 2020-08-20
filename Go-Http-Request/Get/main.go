package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	GetBody()
}

func GetBody() {
	var url string = "https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
