// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package main

import (
	"log"

	"github.com/ALiwoto/argparser/argparser"
)

const boolTest = "/test --flag1 true --flag2 on" +
	" --flag3 false" +
	" --flag4 NO" +
	" --flag5 off"

func main() {
	args, err := argparser.ParseArgDefault(boolTest)
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

	s1 := args.GetAsString("flag4")
	log.Println(s1)

	// lets try more than one flag this time.
	// this method will search for the flags you passed to it.
	// if the first flag doesn't exist, it will look
	// for another one.
	// if there are no such flags at all, it will return you
	// `0, false`
	// if the first flag exists but it can't be
	// converted to integer, it will give you `0, false`s
	i1, ok := args.GetAsInteger("flag6", "flag10")
	if !ok {
		log.Println("couldn't parse flag value to integer")
	} else {
		log.Println("oh yeah, the integer value is ", i1)
	}

}
