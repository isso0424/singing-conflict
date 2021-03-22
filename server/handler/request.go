package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/singing-conflict/server/key"
	"log"
	"net/http"
	"time"
)

const (
	url = "https://api.github.com"
	fetchEndpoint = "/repos/%s/%s/pulls/%d"
	sendEndpoint = "/repos/%s/%s/pulls/%d/comments"
)

func Request(targetRepo string, owner string, number int) {
	go func() {
		for {
			token := key.Generate()
			log.Println(token)
			d, err := fetchPull(targetRepo, owner, number, token)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(d)
			if d.MergeableState == "unknown" || d.MergeableState == "" {
				time.Sleep(time.Second * 5)
			} else if d.MergeableState == "clean" {
				return
			} else {
				fmt.Println("dirty")
				break
			}
		}
	}()
}

type repoData struct {
	MergeableState string `json:"mergeable_state"`
}

func fetchPull(targetRepo string, owner string, number int, token string) (d repoData, err error) {
	req, err := http.NewRequest("GET", url + fmt.Sprintf(fetchEndpoint, owner, targetRepo, number), nil)
	if err != nil {
		return
	}
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

	var da map[string]interface{}
	err = json.Unmarshal(data, &da)
	fmt.Println(da["mergeable_state"])
	fmt.Printf("%v\n", d)

	return
}

func commentPull(targetRepo, owner string, number int, token string) {
}
