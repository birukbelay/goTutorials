package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancle := context.WithCancel(ctx)

	time.AfterFunc(time.Second, cancle) // same code as bellow

	// go func(){
	// 	// s:= bufio.NewScanner(os.Stdin)
	// 	// s.Scan()
	// 	timeSleep(time.Second) //cancle after one second
	// 	cancle()
	// }()

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err(), "err")
	}

}

// same code as above
// func main(){
// 	ctx:= context.Background()
// 	ctx, cancle:= context.WithTimeout(ctx, time.Second)
// 	defer cancel()

// 	sleepAndTalk(ctx, 5*time.Second, "hello")
// }
