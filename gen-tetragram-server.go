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
	monograms := []string{"------", "--  --", "- - - "}

	// TODO: figure out what the bigram chart number is from this also

	// Print the time
	t := time.Now()
	fmt.Fprintf(w, "<html><body><pre>\n")
	fmt.Fprintf(w, "%s\n", t.Format("2006-01-02 15:04:05"))

	// Tetra is four
	for i := 0; i <= 3; i++ {
		// choose random up from [0,3)
		randInt, err := rand.Int(rand.Reader, big.NewInt(3))
		if err != nil {
			panic(err)
		}
		// have to convert bigint back to int64 to use in index
		randMonogram := monograms[randInt.Int64()]
		fmt.Fprintf(w, "%s\n", randMonogram)
	}
	fmt.Fprintf(w, "</pre></body></html>\n")
}

func main() {
	http.HandleFunc("/", generateTetragram)  // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
