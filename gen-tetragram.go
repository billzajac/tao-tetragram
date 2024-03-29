package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// https://en.wiktionary.org/wiki/Taixuanjing
// Tai Xuan Jing tetragrams
//
// There are 81 permutations of the following three lines, given 4 lines
// ------ Heaven, --  -- Earth, - - - Man
// 3^4 = 81

func main() {
	t := time.Now()
	tetragrams := []string{"𝌆", "𝌇", "𝌈", "𝌉", "𝌊", "𝌋", "𝌌", "𝌍", "𝌎", "𝌏", "𝌐", "𝌑", "𝌒", "𝌓", "𝌔", "𝌕", "𝌖", "𝌗", "𝌘", "𝌙", "𝌚", "𝌛", "𝌜", "𝌝", "𝌞", "𝌟", "𝌠", "𝌡", "𝌢", "𝌣", "𝌤", "𝌥", "𝌦", "𝌧", "𝌨", "𝌩", "𝌪", "𝌫", "𝌬", "𝌭", "𝌮", "𝌯", "𝌰", "𝌱", "𝌲", "𝌳", "𝌴", "𝌵", "𝌶", "𝌷", "𝌸", "𝌹", "𝌺", "𝌻", "𝌼", "𝌽", "𝌾", "𝌿", "𝍀", "𝍁", "𝍂", "𝍃", "𝍄", "𝍅", "𝍆", "𝍇", "𝍈", "𝍉", "𝍊", "𝍋", "𝍌", "𝍍", "𝍎", "𝍏", "𝍐", "𝍑", "𝍒", "𝍓", "𝍔", "𝍕", "𝍖"}

	// choose random from [0,81)
	randInt, err := rand.Int(rand.Reader, big.NewInt(81))
	if err != nil {
		panic(err)
	}

	// have to convert bigint back to int64 to use in index
	tetragram := tetragrams[randInt.Int64()]
	tetragram_int := randInt.Int64() + 1
	// now := t.Format("2006-01-02 15:04:05")
	fmt.Printf("%d %s %s\n", tetragram_int, tetragram, t)
}
