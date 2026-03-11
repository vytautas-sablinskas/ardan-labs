package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userInfo struct {
	Name     string
	NumRepos int `json:"public_repos"`
}

const (
	githubUsername = "vytautas-sablinskas"
)

func main() {
	userInfo, err := FetchGithubUserInfo(githubUsername)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("Github name:", userInfo.Name)
	fmt.Println("Public repository count:", userInfo.NumRepos)
}

func FetchGithubUserInfo(username string) (userInfo, error) {
	url := "https://api.github.com/users/"
	resp, err := http.Get(url + username)
	if err != nil {
		return userInfo{}, fmt.Errorf("get: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return userInfo{}, fmt.Errorf("invalid status: %s", resp.Status)
	}

	var userInformation userInfo
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&userInformation); err != nil {
		return userInfo{}, fmt.Errorf("decode: %w", err)
	}

	return userInformation, nil
}
