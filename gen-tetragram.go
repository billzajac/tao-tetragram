package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	monograms := []string{"------", "--  --", "- - - "}

	// Print the time
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	// Tetra is four
	for i := 0; i <= 3; i++ {
		// choose random up from [0,3)
		randInt, err := rand.Int(rand.Reader, big.NewInt(3))
		if err != nil {
			panic(err)
		}
		// have to convert bigint back to int64 to use in index
		randMonogram := monograms[randInt.Int64()]
		fmt.Printf("%s\n", randMonogram)
	}
}
