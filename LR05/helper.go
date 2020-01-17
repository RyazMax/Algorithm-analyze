package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func sleeper(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	profile, err := os.OpenFile("profile.4", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal("Hello World")
	}

	toSleep := 5
	for i := 1; i < 12; i++ {
		begin := time.Now()
		for j := 0; j < i; j++ {
			sleeper(toSleep)
		}
		last := time.Since(begin)
		s := fmt.Sprintf("%f ", last.Seconds())
		profile.WriteString(s)
	}
	profile.Close()
}
