package cookie

import (
	"fmt"
	"log"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}
func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	fmt.Printf("starting...")
	log.Fatal(server.ListenAndServe())
}
