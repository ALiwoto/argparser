// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package argparser

import (
	"errors"
	"strings"

	"github.com/ALiwoto/argparser/interfaces"
	ws "github.com/ALiwoto/argparser/wotoStrings"
	"github.com/ALiwoto/argparser/wotoStrong"
	wv "github.com/ALiwoto/argparser/wotoValues"
)

// Flag is the options passed along with the commands
// by users. they should send them with prefex "--",
// but we will remove them in the argparser.
type Flag struct {
	name  string
	index int
	value interface{}
	fType FlagType
}

type EventArgs struct {
	command string // command without '/' or '!'
	flags   []Flag
	rawData string
}

// ParseArg will parse the whole text into an EventArg and will return it.
func ParseArg(text string) (e *EventArgs, err error) {
	if ws.IsEmpty(&text) {
		return nil, errors.New("text cannot be empty")
	}

	ss := wotoStrong.Ss(text)
	if !ss.HasPrefix(wv.COMMAND_PREFIX1, wv.COMMAND_PREFIX2) {
		return nil, errors.New("this message is not a command at all")
	}

	cmdR := ss.SplitStr(wv.SPACE_VALUE)
	if len(cmdR) == wv.BaseIndex {
		return nil, errors.New("wasn't able to get the command")
	}

	cmd := cmdR[wv.BaseIndex]
	if cmd.IsEmpty() {
		return nil, errors.New("length of the command cannot be zero")
	}

	cmdSs := cmd.TrimStr(wv.COMMAND_PREFIX1, wv.COMMAND_PREFIX2, wv.SUDO_PREFIX1, wv.SPACE_VALUE)
	if cmdSs.IsEmpty() {
		return nil, errors.New("command cannot be only whitespace")
	}

	cmdStr := cmdSs.GetValue()

	e = &EventArgs{
		command: cmdStr,
	}

	// lock the special characters such as "--", ":", "=".
	ss.LockSpecial()

	tmpOSs := ss.SplitStr(wv.FLAG_PREFIX)
	// check if we have any flags or not.
	// I think this if is not necessary actually,
	// but I just added it to prevent some cases of
	// panics. and also it will reduce the time order
	// I guess.
	if len(tmpOSs) < wv.BaseTwoIndex {
		// please notice that we should send the original
		// text to this method.
		// because our locked QString contains JA characters
		// and should not be used here.
		lookRaw(&text, e)
		return e, nil
	}

	flagsR := tmpOSs[wv.BaseOneIndex:]
	// it means it has no flags available.
	// so return the e.
	if len(flagsR) == wv.BaseIndex {
		// please notice that we should send the original
		// text to this method.
		// because our locked QString contains JA characters
		// and should not be used here.
		lookRaw(&text, e)
		return e, nil
	}

	myFlags := make([]Flag, wv.BaseIndex)
	tmp := wv.EMPTY
	var tmpFlag Flag
	var tmpArr []interfaces.QString

	for i, current := range flagsR {
		tmpFlag = Flag{
			index: i,
		}

		tmp = wv.EMPTY
		// let me explain you something here.
		// it really does matter how you pass these constants to this functions.
		// first of all should be equal.
		// and then double dot (':')
		// and in the end, it should be space.
		// please don't forget that you should prioritize them.
		tmpArr = current.SplitStrFirst(wv.EqualStr, wv.DdotSign, wv.SPACE_VALUE)

		tmpFlag.setNameQ(tmpArr[wv.BaseIndex])
		if len(tmpArr) == wv.BaseIndex {
			tmpFlag.setNilValue()
			myFlags = append(myFlags, tmpFlag)
			continue
		}

		for i, ar := range tmpArr {
			if i == wv.BaseIndex {
				// ignore first slice, because it's flag name.
				continue
			}

			ar.UnlockSpecial()
			tmp += ar.GetValue()
		}
		tmpFlag.setRealValue(tmp)

		myFlags = append(myFlags, tmpFlag)

	}

	e.setFlags(myFlags)

	return e, nil
}

// look raw will look for raw data.
// please use this function when and only when
// no flags are provided for our commands.
func lookRaw(text *string, e *EventArgs) {
	owoStr := strings.SplitN(*text, e.command, wv.BaseTwoIndex)
	if len(owoStr) < wv.BaseTwoIndex {
		return
	}

	tmp := strings.Join(owoStr[wv.BaseOneIndex:], wv.EMPTY)
	tmp = strings.TrimSpace(tmp)

	e.rawData = tmp
}
