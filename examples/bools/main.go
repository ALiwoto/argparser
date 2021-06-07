package main

import (
	"log"

	"github.com/ALiwoto/argparser"
)

const boolTest = "/test --flag1 true --flag2 on" +
	" --flag3 false" +
	" --flag4 NO" +
	" --flag5 off"

func main() {
	args, err := argparser.ParseArg(boolTest)
	if err != nil {
		log.Fatal(err)
	}

	if args.HasFlag("flag1", "anotherFlag") {
		log.Println("it has flag1 or anotherFlag")
	}

	b1 := args.GetAsBool("flag1")
	if b1 {
		log.Println("it's true!")
	} else {
		log.Println("it's false!")
	}
}
