package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func askConfirm(t string) (c bool) {
	fmt.Printf("%s: (y/N) ", t)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		if t == "y" {
			c = true
		}
		break
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return c
}

func readInput(t string) (i string) {
	fmt.Printf("%s\n", t)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i = s.Text()
		// r, w := utf8.DecodeRuneInString(i)

		// log.Println(r == '\n')
		// log.Println(w)

		// log.Println(i == "\n")
		// log.Println(i == " ")
		// log.Println(i == "")
		// log.Printf("'%s'", i)
		break
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return i
}
