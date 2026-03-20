package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5000 * time.Millisecond)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Println(t)
	}
}
