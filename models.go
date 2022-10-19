package main

import (
	"fmt"

	"github.com/tonyskapunk/yctf/pkg/flag"
)

// challenge is a specific task to be accomplished, returning a flag once
// completed.
type challenge struct {
	// ID is the internal reference to the challenge's ID.
	ID int
	// Flag is what's captured.
	Flag *flag.Flag
	// Template is what's rendered
	Template string
}

// getFlag returns a flag on the specified index or an error
func getFlag(n int) (challenge, error) {
	var c challenge

	if n < 0 || n > len(challenges) {
		return c, fmt.Errorf("index out of range: %v", n)
	}

	return challenges[n], nil
}
