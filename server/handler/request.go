package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/singing-conflict/server/key"
	"log"
	"net/http"
)

const (
	url = "https://api.github.com"
	fetchEndpoint = "/repos/%s/%s/pulls/%d"
	sendEndpoint = "/repos/%s/%s/pulls/%d/comments"
)

func Request(targetRepo string, owner string, number int) {
	go func() {
		token := key.Generate()
		_, err := fetchPull(targetRepo, owner, number, token)
		if err != nil {
			log.Println(err)
			return
		}
	}()
}

func fetchPull(targetRepo string, owner string, number int, token string) (d map[string]interface{}, err error) {
	req, err := http.NewRequest("GET", url + fmt.Sprintf(fetchEndpoint, owner, targetRepo, number), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "token " + token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	client := http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &d)
	if err != nil {
		return
	}

	fmt.Printf("%v\n", &d)

	return
}
