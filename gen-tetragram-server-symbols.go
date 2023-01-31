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
	t := time.Now()
	tetragrams := []string{"𝌆", "𝌇", "𝌈", "𝌉", "𝌊", "𝌋", "𝌌", "𝌍", "𝌎", "𝌏", "𝌐", "𝌑", "𝌒", "𝌓", "𝌔", "𝌕", "𝌖", "𝌗", "𝌘", "𝌙", "𝌚", "𝌛", "𝌜", "𝌝", "𝌞", "𝌟", "𝌠", "𝌡", "𝌢", "𝌣", "𝌤", "𝌥", "𝌦", "𝌧", "𝌨", "𝌩", "𝌪", "𝌫", "𝌬", "𝌭", "𝌮", "𝌯", "𝌰", "𝌱", "𝌲", "𝌳", "𝌴", "𝌵", "𝌶", "𝌷", "𝌸", "𝌹", "𝌺", "𝌻", "𝌼", "𝌽", "𝌾", "𝌿", "𝍀", "𝍁", "𝍂", "𝍃", "𝍄", "𝍅", "𝍆", "𝍇", "𝍈", "𝍉", "𝍊", "𝍋", "𝍌", "𝍍", "𝍎", "𝍏", "𝍐", "𝍑", "𝍒", "𝍓", "𝍔", "𝍕", "𝍖"}

	// choose random from [0,81)
	randInt, err := rand.Int(rand.Reader, big.NewInt(81))
	if err != nil {
		panic(err)
	}

	// have to convert bigint back to int64 to use in index
	tetragram := tetragrams[randInt.Int64()]
	// Remember to add 1 here because the slice starts with position 0
	tetragram_int := randInt.Int64() + 1
	// now := t.Format("2006-01-02 15:04:05")

	fmt.Fprintf(w, "%d %s %s\n", tetragram_int, tetragram, t)
}

func main() {
	http.HandleFunc("/", generateTetragram)  // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
