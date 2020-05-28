package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mtSigningKey = os.Get("MY_TOKEN_J")
var mySigningKey = []byte("secretDRAVAWordSigning")

func homePage(w http.ResponseWriter, r *http.Request) {

	validToken2, err := GenerateJWT()
	if err != nil {
		fmt.Println("..1")
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", validToken2)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error%s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, "Error%s", err.Error())
	}
	fmt.Fprintf(w, string(body))
}

// GenerateJWT ...
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("..2")

		fmt.Errorf("something went wrong %s", err.Error)
		return "", err
	}
	return tokenString, nil

}

// GenerateJWT2 ...
func GenerateJWT2() (string, error) {

	validToken1, err := GenerateJWT()

	if err != nil {
		fmt.Println("..3")

		fmt.Printf(err.Error())
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(validToken1))

	if err != nil {
		fmt.Println("..4")
		fmt.Println("something went wrong %s", err.Error())

		fmt.Errorf("something went wrong %s", err.Error)
		return "", err
	}
	return tokenString, nil

}

func handlers() {
	http.HandleFunc("/", homePage)
	fmt.Println("...9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {

	fmt.Println("my simple client")
	handlers()

}
