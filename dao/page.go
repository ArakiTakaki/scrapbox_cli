package dao

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

type Page struct {
  ProjectName string    `json:"projectName"`
  Skip        int       `json:"skip"`
  Count       int       `json:"count"`
  Pages       []Article `json:"pages"`
}
type User struct {
  ID string `json:"id"`
}
type Article struct {
  Id                string   `json:"id"`
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

func Say(article string, c chan Page) {
  url := "https://scrapbox.io/api/pages/" + article + "/"
  resp, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)

  var page Page

  if err := json.Unmarshal(byteArray, &page); err != nil {
    panic(err)
  }
  c <- page
}
