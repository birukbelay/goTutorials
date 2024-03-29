package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secretDRAVAWordSigning")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there waw error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "not autorized")
		}
	})
}

func HandleRequest() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("server")
	HandleRequest()
}
