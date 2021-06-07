// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package argparser

import (
	"strconv"
	"strings"

	tf "github.com/ALiwoto/argparser/interfaces"
	wv "github.com/ALiwoto/argparser/wotoValues"
)

//---------------------------------------------------------

func (f *Flag) setName(name string) {
	name = strings.Trim(name, wv.SPACE_VALUE)
	f.name = name
}

func (f *Flag) setNameQ(q tf.QString) {
	f.setName(q.GetValue())
}

func (f *Flag) setNilValue() {
	f.value = nil
	f.fType = NoneFlagType
}

func (f *Flag) setRealValue(value string) {
	myI, err := strconv.ParseInt(value, wv.BaseTen, wv.Base64Bit)
	if err == nil {
		f.setAsInt(&myI)
		return
	}

	tmp := strings.Trim(value, wv.SPACE_VALUE)
	if len(tmp) == wv.BaseIndex {
		f.setAsBool(true)
		return
	}

	if tmp[wv.BaseIndex] == wv.CHAR_STR {
		myInt := len(tmp) - wv.BaseOneIndex
		if tmp[myInt] == wv.CHAR_STR {
			tmp = strings.TrimPrefix(tmp, wv.STR_SIGN)
			tmp = strings.TrimSuffix(tmp, wv.STR_SIGN)
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
