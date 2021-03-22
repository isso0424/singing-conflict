package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CommentToPR(targetRepo, owner string, number int, token string) (err error) {
	body, err := json.Marshal(map[string]string{
		"body": "conflict歌います。ズォールヒ～～↑ｗｗｗｗヴィヤーンタースｗｗｗｗｗワース フェスツｗｗｗｗｗｗｗルオルｗｗｗｗｗプローイユクｗｗｗｗｗｗｗダルフェ スォーイヴォーｗｗｗｗｗスウェンネｗｗｗｗヤットゥ ヴ ヒェンヴガｒジョｊゴアｊガオガオッガｗｗｗじゃｇｊｊ",
	})
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url+fmt.Sprintf(sendEndpoint, owner, targetRepo, number), bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	log.Println(string(data))

	return
}
