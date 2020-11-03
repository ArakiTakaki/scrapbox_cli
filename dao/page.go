package dao

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Page TODO bookに治す
type Page struct {
	ProjectName string    `json:"projectName"`
	Skip        int       `json:"skip"`
	Count       int       `json:"count"`
	Pages       []Article `json:"pages"`
}

// User TODO bookに治す
type User struct {
	ID string `json:"id"`
}

// Article TODO bookに治す
type Article struct {
	ID                string   `json:"id"`
	Title             string   `json:"title"`
	Image             string   `json:"image"`
	Descriptions      []string `json:"descriptions"`
	User              User     `json:"user"`
	Pin               int      `json:"pin"`
	Views             int      `json:"viwes"`
	CommitID          string   `json:"commitId"`
	CreatedAt         int      `json:"created"`
	UpdatedAt         int      `json:"updated"`
	AccessedAt        int      `json:"accessed"`
	SnapshotCreatedAt int      `json:"snapshotCreated"`
}

// Say hogehoge
func Say(book string) (Page, error) {
	url := "https://scrapbox.io/api/pages/" + book + "/"

	resp, err := http.Get(url)
	if err != nil {
		return Page{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	var page Page

	if err := json.Unmarshal(byteArray, &page); err != nil {
		panic(err)
	}

	return page, nil
}
