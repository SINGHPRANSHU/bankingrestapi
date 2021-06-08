package main

import (
	"fmt"
	"net/http"
	"os"
	"restapi/helper"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func MiddlewareValidateUser(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(r.Header["Authorization"][0], func(t *jwt.Token) (interface{}, error) {
               if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				   return nil, fmt.Errorf("there is an error")
			   }
			   return []byte(os.Getenv("secret")), nil
			})
			if err!= nil {
				helper.Insufficient("not authorized",w)
				return
			}

			if token.Valid {
				if token.Claims.(jwt.MapClaims)["expiry"].(float64) < float64(time.Now().Unix()) {
                   helper.Insufficient("token expired", w)
				   return 
				}
				r.Header.Add("userid",token.Claims.(jwt.MapClaims)["user"].(string))
				next.ServeHTTP(w, r)
				return
			}
			
			
		}
		helper.Insufficient("not authorized",w)
		// call the next handler
		
	})
}