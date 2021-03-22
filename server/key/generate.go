package key

import (
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

	return tokenString
}
