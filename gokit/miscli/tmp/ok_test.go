package tmp

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"testing"
	"time"
)

func TestTmp(t *testing.T) {
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()
	for {
		if err := r.WaitN(ctx, 2); err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Now().Format(time.RFC3339))
		time.Sleep(time.Second)
	}
}
func TestTmp2(t *testing.T) {
	r := rate.NewLimiter(1, 5)
	for {
		if cando := r.AllowN(time.Now(),2); cando {
			fmt.Println(time.Now().Format(time.RFC3339))
		} else {
			fmt.Println("too many request")
		}
		time.Sleep(time.Second)
	}
}
