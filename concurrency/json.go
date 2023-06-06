package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Declare a variable 'client' as an instance of http.Client
var client = &http.Client{Timeout: 10 * time.Second}

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type RamdonUser struct {
	Results []UserResult
}

type UserResult struct {
	Name UserName `json:"name"`
	Email string
	Picture UserPicture
}

type UserName struct {
	Title string
	First string
	Last string
}

type UserPicture struct {
	Large string
	Medium string
	Thumbnail string
}

func GetCatFact() {
	url := "https://catfact.ninja/fact"

	var catFact CatFact

	err := GetJson(url, &catFact) // Pass the address of catFact to GetJson
	if err != nil {
		fmt.Printf("error getting cat fact: %s\n", err.Error())
		return
	} else {
		fmt.Printf("This is a Cat fact: %s\n", catFact.Fact)
	}
}

func GetRandomUSer() {
	url := "https://randomuser.me/api?inc=name,email,picture"

	var user RamdonUser

	err := GetJson(url, &user)

	if err != nil {
		fmt.Printf("error getting json: %s\n", err.Error())
	} else {
		fmt.Printf("User: %s %s %s\nEmail: %s\nThumbnail: %s",
			user.Results[0].Name.Title,
			user.Results[0].Name.First,
			user.Results[0].Name.Last,
			user.Results[0].Email,
			user.Results[0].Picture.Thumbnail,

		)
	}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	GetCatFact()
	GetRandomUSer()
}
