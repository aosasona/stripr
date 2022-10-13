package main

import "os"

func main() {
	s := Scanner{}
	s.New(os.Args).Run()
}
