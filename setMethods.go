// Bot.go Project
// Copyright (C) 2021 Sayan Biswas, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package argparser

import (
	"strconv"
	"strings"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
)

//---------------------------------------------------------

func (f *Flag) setName(name string) {
	name = strings.Trim(name, ws.SPACE_VALUE)
	f.name = name
}

func (f *Flag) setNameQ(q ws.QString) {
	f.setName(q.GetValue())
}

func (f *Flag) setNilValue() {
	f.value = nil
	f.fType = NoneFlagType
}

func (f *Flag) setRealValue(value string) {
	myI, err := strconv.ParseInt(value, ws.BaseTen, ws.Base64Bit)
	if err == nil {
		f.setAsInt(&myI)
		return
	}

	tmp := strings.Trim(value, ws.SPACE_VALUE)
	if len(tmp) == ws.BaseIndex {
		f.setAsBool(true)
		return
	}

	if tmp[ws.BaseIndex] == ws.CHAR_STR {
		myInt := len(tmp) - ws.BaseOneIndex
		if tmp[myInt] == ws.CHAR_STR {
			tmp = strings.TrimPrefix(tmp, ws.STR_SIGN)
			tmp = strings.TrimSuffix(tmp, ws.STR_SIGN)
			f.setAsString(&tmp)
			return
		}
	}

	v, isBool := ToBoolType(value)
	if isBool {
		f.setAsBool(v)
		return
	}

	f.setAsString(&value)
}

func (f *Flag) setAsBool(value bool) {
	f.value = value
	f.fType = BoolFlagType
}

func (f *Flag) setAsString(value *string) {
	f.value = *value
	f.fType = StringFlagType
}

func (f *Flag) setAsInt(value *int64) {
	f.value = *value
	f.fType = Int64FlagType
}

//---------------------------------------------------------

func (e *EventArgs) setFlags(f []Flag) {
	e.flags = f
}

//---------------------------------------------------------
