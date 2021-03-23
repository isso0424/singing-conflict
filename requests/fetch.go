package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type repoData struct {
	MergeableState string `json:"mergeable_state"`
}

func FetchPP(targetRepo string, owner string, number int) (d repoData, err error) {
	req, err := http.NewRequest("GET", url+fmt.Sprintf(fetchEndpoint, owner, targetRepo, number), nil)
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

	return
}
