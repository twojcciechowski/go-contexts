package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func main() {

	ctx := context.WithValue(context.Background(), "Key1", "Value1" )
	ctx,cancelFunc := context.WithTimeout(ctx, 3*time.Second )

	defer cancelFunc()
	out:= make(chan string)
	runWithTimeout(ctx, out)

	select {

		case <- ctx.Done(): {
				log.Printf("Processing to long \n")
				log.Print(ctx.Err())
		}
		case s := <- out: {
			log.Printf("Processing ended with result: %s \n", s);
		}
	}
}

func runWithTimeout(ctx context.Context, out chan string){
	log.Printf("Starting with value Key1 = %v\n",ctx.Value("Key1"))

	go func() {
		log.Printf("Staring processing\n")

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		sec := r1.Float64() * 5

		duration:= time.Duration(sec) * time.Second
		log.Printf("Processing will take : %v , %v \n", sec, duration)
		time.Sleep( duration)
		log.Printf("Processing finished\n")
		out <- "Ready"
	}()

}