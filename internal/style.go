package internal 

import ( 
	"fmt"
	"time"
	"math/rand"
)

func TypePrint(s string) {
	for _, char := range s {
		fmt.Print(string(char))

		min := 20
		max := 130
		delay := time.Duration(rand.Intn(max-min)+min) * time.Millisecond
		time.Sleep(delay)
	}
	fmt.Print("\n")
}
