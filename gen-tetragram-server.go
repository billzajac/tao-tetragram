package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
)

func generateTetragram(w http.ResponseWriter, r *http.Request) {
	base10Sum := 0
	monograms := []string{"------", "--  --", "- - - "}
	// 0000 -> 1
	// 0001 -> 2
	// 0002 -> 3
	// 0010 -> 4
	// 0011 -> 5
	// 0012 -> 6
	// 0020 -> 7
	// 0021 -> 8
	// 0022 -> 9
	// 0100 -> 10
	// 0101 -> 11

	// TODO: figure out what the bigram chart number is from this also

	// Print the time
	t := time.Now()
	fmt.Fprintf(w, "<html><body><pre>\n")
	fmt.Fprintf(w, "%s\n", t.Format("2006-01-02 15:04:05"))

	// Tetra is four
	for i := 0; i <= 3; i++ {
		// choose random up from [0,3)
		randBigInt, err := rand.Int(rand.Reader, big.NewInt(3))
		if err != nil {
			panic(err)
		}
		// have to convert bigint back to int64 to use in index
		randInt := int(randBigInt.Int64())

		// To convert from trinary/ternary (base 3) to base 10, multiply each part by it's order of magnitude
		switch i {
		case 0:
			base10Sum = base10Sum + randInt*27
		case 1:
			base10Sum = base10Sum + randInt*9
		case 2:
			base10Sum = base10Sum + randInt*3
		case 3:
			base10Sum = base10Sum + randInt*1
		}

		// have to convert bigint back to int64 to use in index
		randMonogram := monograms[randBigInt.Int64()]
		fmt.Fprintf(w, "%s\n", randMonogram)
	}
	// Remember to add 1 here because this system starts with 1, not 0
	fmt.Fprintf(w, "\nPassage: %d\n", base10Sum+1)
	fmt.Fprintf(w, "</pre></body></html>\n")
}

func main() {
	http.HandleFunc("/", generateTetragram)  // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
