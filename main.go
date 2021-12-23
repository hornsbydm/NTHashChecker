package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

		r.IsCompromised = r.isCompromised(queryHash(r.NTHash))

		fmt.Printf("User: %v\tIsCompromised: %v\n", r.User, r.IsCompromised)

	}

}
