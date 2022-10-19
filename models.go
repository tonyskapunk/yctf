package main

import (
	"fmt"

	"github.com/tonyskapunk/yctf/pkg/flag"
)

// type flag struct {
// 	ID       int
// 	Flag     string
// 	Template string
// }
//
// // getAllFlags gets all the flags from a data source
// func getAllFlags() []flag {
// 	var flags []flag
//
// 	data, err := ioutil.ReadFile("./flags.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	err = json.Unmarshal([]byte(data), &flags)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return flags
// }
//
// // getFlag returns a flag on the specified index or an error
// func getFlag(n int) (flag, error) {
// 	var f flag
// 	flags := getAllFlags()
//
// 	if n < 0 || n > len(flags) {
// 		return f, fmt.Errorf("Index out of range: %v", n)
// 	}
//
// 	return flags[n], nil
// }

type flagConstruct struct {
	ID       int
	Flag     string
	Template string
}

// get Flag returns a generated flag from a seed
func getFlag(i int) (flagConstruct, error) {
	seed := "my-random-seed-from-session"

	fact := flag.NewFactory(seed)
	f := fact.NewFlag([]byte(fmt.Sprintf("%d", i)))
	t := fmt.Sprintf("flag%d.html", i)
	var fc = flagConstruct{ID: i, Flag: f.String(), Template: t}

	return fc, nil
}
