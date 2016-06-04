package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:\n%s {path}\n", os.Args[0])
		os.Exit(-1)
	}
	path := os.Args[1]
	sparses, err := FindSparses(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, sparse := range sparses {
		fmt.Println(sparse)
	}
}
