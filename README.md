# Arg Parser
====================================



A rich tool for parsing flags and values in pure Golang.
No additional library is required and you can use everywhere.



## Features

* Supporting different formats:
```go
	int (int64, int32, int16, int8)
	uint (uint64, uint32, uint16, uint8)
	string // (you can easily convert any type to string)
	bool // yes, on, true, 1 = true; no, off, false, 0 = false
```

* Easily convert each type to another types. with only one functions.

* Low time order! It's as fast as spliting two strings and then joining them together XD

* Simple algorithms. You can easily read the algorithms used in this library, and you are allowed to copy and use them in your own projects (even if they are in Golang or not)!

* No additional libs. You don't have to waste your time and storage for downloading another libraries and dependencies. Only this lib, will meet your needs.

* Supports multi flags with the same name (Of course they have different indexes).

* Rigorously tested. We are testing the lib in any possible ways. If you find out a bug, please create an issue on our [issue tracker](https://github.com/aliwoto/argparser/issues). We will check it and solve the problem

<hr/>

## Examples:

### parsing a command is a simple work!

```go

const text = "/test --flag1 true --flag2 on" +
	" --flag3 false" +
	" --flag4 \"Hello, how are you?\"" +
	" --flag5 off\n\n\n\n\n\n" +
	" --flag6 123456789"

	//-----------------------------------------------------

	// parse it and get an *EventArgs value.
	args, err := argparser.ParseArg(boolTest)
	if err != nil {
		// if the format of the text is not correct,
		// you will get an error.
		// if you find out any bug here, please report us.
		// if the text doesn't contain ant command, you will
		// get error (it should start with command.)
		log.Fatal(err)
	}

	// check if it contains any flag with name `flag1` or not?
	if args.HasFlag("flag1", "anotherFlag") {
		// doesn't matter this
		log.Println("it has flag1 or anotherFlag")
	}

	// get a flag's value as bool.
	// if the flags are not present in EventArgs at all,
	// this method will return false.
	// if it's on, yes, true, 1 (one integer), it returns true
	// if it's off, no, false, 0, it returns false
	// if it's string, it will return true
	// if it doesn't have any value (empty string), it returns true
	// 
	b1 := args.GetAsBool("flag1")
	if b1 {
		log.Println("it's true!")
	} else {
		log.Println("it's false!")
	}

	s1 := args.GetAsString("flag4")
	log.Println(s1) // Hello, how are you?

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

```

> need more examples?
 then take a look at [example directory](examples/)!

<hr/>

## How to get started?

* Download the library with the standard go get command:

`go get github.com/ALiwoto/argparser`


## Support and Contributions

If you think you have found a bug or have a feature request, feel free to use our [issue tracker](https://github.com/ALiwoto/argparser/issues). Before opening a new issue, please search to see if your problem has already been reported or not.  Try to be as detailed as possible in your issue reports.

If you need help using argparser or have other questions we suggest you to join our [telegram community](https://t.me/chiruzon).  Please do not use the GitHub issue tracker for personal support requests.


## Docs

Docs can be found [here](https://pkg.go.dev/github.com/ALiwoto/argparser).
> Be sure to read them carefuly!

## License

The argparser project is under the [MIT License](https://opensource.org/licenses/MIT).
See the [LICENSE](LICENSE) file for more details.