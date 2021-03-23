package handler

// https://developer.github.com/webhooks/
import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HookContext struct {
	Event     string
	Id        string
	Payload   []byte
}

func ParseHook(secret []byte, req *http.Request) (*HookContext, error) {
	hc := HookContext{}

	if hc.Event = req.Header.Get("x-github-event"); len(hc.Event) == 0 {
		return nil, errors.New("No event!")
	}

	if hc.Id = req.Header.Get("x-github-delivery"); len(hc.Id) == 0 {
		return nil, errors.New("No event Id!")
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	hc.Payload = body

	return &hc, nil
}

type response struct {
	Number      int `json:"number"`
	PullRequest struct {
		Head struct {
			Repo struct {
				Name  string
				Owner struct {
					Login string `json:"login"`
				} `json:"owner"`
			} `json:"repo"`
		} `json:"head"`
	} `json:"pull_request"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("x-gitHub-event") != "pull_request" {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "{}")

		return
	}

	hc, err := ParseHook([]byte(os.Getenv("WEBHOOK_SECRET")), r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook! ('%s')", err)
		io.WriteString(w, "{}")
	}

	var data response
	err = json.Unmarshal(hc.Payload, &data)

	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook! ('%s')", err)
		io.WriteString(w, "{}")
		return
	}

	log.Printf("Received %s", hc.Event)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")
	Request(data.PullRequest.Head.Repo.Name, data.PullRequest.Head.Repo.Owner.Login, data.Number)
	return
}
