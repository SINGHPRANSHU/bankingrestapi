package main

import (
	"fmt"
	"net/http"
	"restapi/helper"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func MiddlewareValidateUser(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
               if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				   return nil, fmt.Errorf("there is an error")
			   }
			   return []byte("secret"), nil
			})
			if err!= nil {
				helper.Insufficient("not authorized",w)
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
		// fmt.Fprintf(w, "not authorized")
		return
		// call the next handler
		
	})
}