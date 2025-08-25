package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserInfo struct {
	Name            string `json:"name"`
	PublicRepoCount int    `json:"public_repos"`
}

func main() {
	userInfo, err := getGithubUserInfo("ardanlabs")
	if err != nil {
		fmt.Println("err fetching github user %w", err)
		return
	}

	fmt.Println("Name:", userInfo.Name)
	fmt.Println("Public repo count:", userInfo.PublicRepoCount)
}

func getGithubUserInfo(username string) (UserInfo, error) {
	url := "https://api.github.com/users/" + username

	resp, err := http.Get(url)
	if err != nil {
		return UserInfo{}, fmt.Errorf("get github info %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("status code %w", err)
	}

	var userInfo UserInfo
	parseResponse(resp.Body, &userInfo)

	return userInfo, nil
}

func parseResponse(body io.ReadCloser, response any) error {
	defer body.Close()

	dec := json.NewDecoder(body)
	if err := dec.Decode(response); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}
