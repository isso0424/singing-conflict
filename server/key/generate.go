package key

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Generate() string {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = 106343
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	key, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("KEY")))

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
		return ""
	}

	request, err := http.NewRequest("GET", "https://api.github.com/app/installations/15627679/access_tokens", nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	request.Header.Add("Authorization", "Bearer " + tokenString)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return ""
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	type d struct {
		Token string `json:"token"`
	}
	log.Println(string(data))
	var result d
	err = json.Unmarshal(data, &result)

	return result.Token
}
