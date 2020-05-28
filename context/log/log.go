package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type Key int

const requestIDKey = Key(42)

// Println ...
func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("cant find id context")
		return
	}
	log.Printf("[%d} %s", id, msg)
}

// Decorate generates new req id for each req
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		// use the new context inside of a request to call th preveous handler func
		// recieving a context , adding a value , sending it back
		f(w, r.WithContext(ctx))
	}
}
