package main

import "log"

var bump = &Bump{}

func main() {
	err := promptAction(bump)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print(bump)
}
