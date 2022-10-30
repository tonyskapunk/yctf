package main

import (
	"fmt"

	"github.com/tonyskapunk/yctf/pkg/flag"
)

type flagConstruct struct {
	ID       int
	Flag     string
	Template string
}

// get Flag returns a generated flag from a seed
func getFlag(i int) (flagConstruct, error) {
	// TODO: Use the session as the seed
	seed := "my-random-seed-from-session"

	fact := flag.NewFactory(seed)
	f := fact.NewFlag([]byte(fmt.Sprintf("%d", i)))
	t := fmt.Sprintf("flag%d.html", i)
	var fc = flagConstruct{ID: i, Flag: f.String(), Template: t}

	return fc, nil
}
