package handler

import (
	"isso0424/singing-conflict/server/key"
	"isso0424/singing-conflict/server/requests"
	"log"
	"time"
)

func Request(targetRepo string, owner string, number int) {
	go func() {
		for {
			d, err := requests.FetchPP(targetRepo, owner, number)
			if err != nil {
				log.Println(err)
				return
			}

			if d.MergeableState == "unknown" || d.MergeableState == "" {
				time.Sleep(time.Second * 5)
				continue
			}

			if d.MergeableState == "clean" {
				return
			}

			token, err := key.Generate()
			if err != nil {
				log.Println(err)
			}

			err = requests.CommentToPR(targetRepo, owner, number, token)
			if err != nil {
				log.Println(err)
			}

			return
		}
	}()
}
