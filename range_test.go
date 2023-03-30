package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestRange(t *testing.T) {
	counts := map[string]int{}

	sentence := "The quick brown fox jumps over the lazy dog"

	words := strings.Fields(strings.ToLower(sentence))

	for _, w := range words {
		// if the word is already in the map, increment it
		// otherwise, set it to 1 and add it to the map
		counts[w]++
	}
	fmt.Println("------\n", counts, "\n------")
	key := "fox"
	_, exists := counts[key]
	if exists {
		fmt.Printf("%q key exists in the map\n", key)
	}
}

func TestMonths(t *testing.T) {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	keys := make([]int, 0, len(months))
	fmt.Println("-------------------")
	for key, month := range months {
		keys = append(keys, key)
		fmt.Println(month)
	}
	sort.Ints(keys)
	fmt.Println("-------------------")
	for _, k := range keys {
		fmt.Println(months[k])
	}
	fmt.Println("-------------------")
}

func TestJSON(t *testing.T) {
	type User struct {
		ID       int
		Name     string
		Phone    string
		Password string
	}
	u := User{
		ID:       1,
		Name:     "Amy",
		Password: "goIsAwesome",
	}

	// create a new JSON encoder
	// that will write to stdout
	enc := json.NewEncoder(os.Stdout)

	// encode the user
	if err := enc.Encode(u); err != nil {

		// handle an error if one occurs
		log.Fatal(err)
	}
}
