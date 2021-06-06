package helper

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateToken(userId primitive.ObjectID)(string,error) {
	Mysigningkey := []byte(os.Getenv("secret"))
    token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Authorized"] = true
	claims["user"] = userId
	claims["expiry"] = time.Now().Add(time.Minute * 720).Unix()

	tokenString, err := token.SignedString(Mysigningkey)
	if err!=nil {
		return "", err
	}
	return tokenString,nil
}