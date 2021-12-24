package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	f := flag.String("f", "", "dump file")
	flag.Parse()

	if *f == "" {
		panic("No File")
	}

	file, err := os.Open(*f)

	if err != nil {
		panic("Can't Open file")
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {

		r := NewRecord(s.Text())

		r.IsCompromised = strings.Contains(queryHash(r.NTHash), r.NTHash[5:])

		fmt.Printf("User: %v\tIsCompromised: %v\n", r.User, r.IsCompromised)

	}

}
