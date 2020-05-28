package main

import (
	"crypto/rand"
	"fmt"

	oklogULID "github.com/oklog/ulid"
)

var ULIDoklog = oklogULID.MustNew(oklogULID.Now(), rand.Reader)
var a = 6

func main() {

	fmt.Println(a)
	fmt.Println(ULIDoklog)
}

// func genUlid() {
// 	t := time.Now().UTC()
// 	entropy := rand.New(rand.NewSource(t.UnixNano()))
// 	id := ulid.MustNew(ulid.Timestamp(t), entropy)
// 	fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
// }

// func genSonyflake() {
// 	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
// 	id, err := flake.NextID()
// 	if err != nil {
// 		log.Fatalf("flake.NextID() failed with %s\n", err)
// 	}
// 	// Note: this is base16, could shorten by encoding as base62 string
// 	fmt.Printf("github.com/sony/sonyflake:      %x\n", id)
// }
