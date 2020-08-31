package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type flag struct {
	ID       int
	Flag     string
	Template string
}

// getAllFlags gets all the flags from a data source
func getAllFlags() []flag {
	var flags []flag

	data, err := ioutil.ReadFile("./flags.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(data), &flags)
	if err != nil {
		log.Fatal(err)
	}

	return flags
}

// getFlag returns a flag on the specified index or an error
func getFlag(n int) (flag, error) {
	var f flag
	flags := getAllFlags()

	if n < 0 || n > len(flags) {
		return f, fmt.Errorf("Index out of range: %v", n)
	}

	return flags[n], nil
}
