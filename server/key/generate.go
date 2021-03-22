package key

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Generate() (webToken string, err error) {
	id, err := strconv.Atoi(os.Getenv("APP_ID"))
	if err != nil {
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	key, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("KEY")))

	tokenString, err := token.SignedString(key)
	if err != nil {
		return
	}

	request, err := http.NewRequest("POST", "https://api.github.com/app/installations/15627679/access_tokens", nil)
	if err != nil {
		return
	}
	request.Header.Add("Authorization", "Bearer "+tokenString)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	type d struct {
		Token string `json:"token"`
	}

	var result d
	err = json.Unmarshal(data, &result)
	webToken = result.Token

	return
}
