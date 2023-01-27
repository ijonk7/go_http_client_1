package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Post struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int64  `json:"userId"`
}

// Index (Contoh Request GET)
func Index(w http.ResponseWriter, r *http.Request) {
	c := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}
	resp, err := c.Get("https://go.dev/")

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf("Body : %s", body)
}

// Store (Contoh Request POST)
func Store(w http.ResponseWriter, r *http.Request) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	postData := bytes.NewBuffer([]byte(`{"post":"boom boom library"}`))
	resp, err := c.Post("https://go.dev/", "application/json", postData)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf("Body : %s", body)
}

// Store2 HTTP Request Dengan Form Data
var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
