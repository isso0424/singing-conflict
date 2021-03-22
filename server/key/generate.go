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
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "Iv1.f77d0d05fba5649c"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, _ := token.SignedString([]byte(os.Getenv("KEY")))

	request, err := http.NewRequest("POST", "https://api.github.com/app/installations/1/access_tokens", nil)
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
	var result d
	err = json.Unmarshal(data, &result)

	return result.Token
}
