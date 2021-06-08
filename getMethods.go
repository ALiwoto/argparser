// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package argparser

import (
	"strconv"
	"strings"

	"github.com/ALiwoto/argparser/wotoStrings"
	wv "github.com/ALiwoto/argparser/wotoValues"
)

//---------------------------------------------------------

// GetValue will return you the value of this flag.
// remember that value can be
func (f *Flag) GetValue() interface{} {
	return f.value
}

// GetType will return the type of the value of this
// flag. please notice that if you don't enter
// any value for this flag, it will be considered as
// `true` (which is type bool).
func (f *Flag) GetType() FlagType {
	return f.fType
}

// GetIndex will return the index of this flag.
func (f *Flag) GetIndex() int {
	return f.index
}

// GetName will give you the name of this flag.
func (f *Flag) GetName() string {
	return f.name
}

func (f *Flag) GetValueAndType() (interface{}, FlagType) {
	return f.GetValue(), f.GetType()
}

// GetAsString will give you the value as an string.
// for example if value is `true` (with flag type of
// `bool`), then the string will be "true".
// or if it's 10(an integer), it will give you: "10".
func (f *Flag) GetAsString() string {
	v, t := f.GetValueAndType()
	if t == StringFlagType {
		vS, ok := v.(string)
		if !ok {
			return NoneTypeStr
		}

		return vS
	}

	if t == BoolFlagType {
		vB, ok := v.(bool)
		if !ok {
			return NoneTypeStr
		}

		if vB {
			return TrueHlc
		} else {
			return FalseHlc
		}
	}

	if t.IsInteger() {
		vI, ok := v.(int64)
		if !ok {
			return NoneTypeStr
		}

		return strconv.FormatInt(vI, wv.BaseTen)
	}

	return NoneTypeStr
}

// GetAsInteger will give the value of this flag as an
// int64 value.
// please notice that if it fails to convert the value
// to int64, it will return you zero and the second
// return value will be false.
func (f *Flag) GetAsInteger() (vI int64, ok bool) {
	v, t := f.GetValueAndType()
	if t == StringFlagType {
		vS, ok := v.(string)
		if !ok {
			return wv.BaseIndex, false
		}

		vS = strings.ReplaceAll(vS, wv.SPACE_VALUE, wv.EMPTY)
		vS = strings.ReplaceAll(vS, wv.STR_SIGN, wv.EMPTY)
		vI, err := strconv.ParseInt(vS, wv.BaseTen, wv.Base64Bit)
		if err != nil {
			return wv.BaseIndex, false
		}

		return vI, true
	}

	if t == BoolFlagType {
		vB, ok := v.(bool)
		if !ok {
			return wv.BaseIndex, false
		}

		if vB {
			return wv.BaseOneIndex, true
		} else {
			return wv.BaseIndex, true
		}
	}

	if t.IsInteger() {
		vI, ok := v.(int64)
		if !ok {
			return wv.BaseIndex, false
		}

		return vI, true
	}

	return wv.BaseIndex, false
}

func (f *Flag) GetAsBool() bool {
	v, t := f.GetValueAndType()
	if t == StringFlagType {
		vS, ok := v.(string)
		if !ok {
			return false
		}

		vS = strings.ReplaceAll(vS, wv.SPACE_VALUE, wv.EMPTY)
		vS = strings.ReplaceAll(vS, wv.STR_SIGN, wv.EMPTY)
		vB, ok := ToBoolType(vS)
		if !ok {
			return false
		}

		return vB
	}

	if t == BoolFlagType {
		vB, ok := v.(bool)
		if !ok {
			return false
		}

		return vB
	}

	if t.IsInteger() {
		vI, ok := v.(int64)
		if !ok {
			return false
		}

		return vI != wv.BaseIndex
	}

	return false
}

//---------------------------------------------------------
func (e *EventArgs) GetCommand() string {
	return e.command
}

func (e *EventArgs) GetFlags() []Flag {
	return e.flags
}

// HasFlag will check if this EventArgs has at least
// one the provided flags or not.
// please notice that if you want to check if it has
// ALL of the flags, use `HasFlags` method.
func (e *EventArgs) HasFlag(names ...string) bool {
	if names == nil {
		return false
	}

	for _, current := range e.flags {
		for _, name := range names {
			if strings.EqualFold(current.name, name) {
				return true
			}
		}
	}

	return false
}

// HasFlag will check if this EventArgs has
// ALL of the provided flags or not.
// please notice that if you want to check if it has
// at least one of the flags, use `HasFlag` method.
func (e *EventArgs) HasFlags(names ...string) bool {
	for _, current := range e.flags {
		for _, name := range names {
			if !strings.EqualFold(current.name, name) {
				return false
			}
		}
	}

	return true
}

func (e *EventArgs) GetIndexFlag(index int) *Flag {
	if index < wv.BaseIndex || index >= e.GetLength() {
		return nil
	}

	return &e.flags[index]
}

func (e *EventArgs) GetFlag(names ...string) *Flag {
	for _, current := range e.flags {
		for _, name := range names {
			if strings.EqualFold(current.name, name) {
				return &current
			}
		}
	}

	return nil
}

func (e *EventArgs) GetLength() int {
	return len(e.flags)
}

func (e *EventArgs) IsEmpty() bool {
	return len(e.flags) == wv.BaseIndex
}

func (e *EventArgs) IsEmptyOrRaw() bool {
	return e.IsEmpty() && len(e.rawData) == wv.BaseIndex
}

// GetAsString will give you the string value of the flag
// with the specified name.
// if there is no flag with this name, or there is an
// error on our path, it will return you an empty string.
func (e *EventArgs) GetAsString(name ...string) string {
	f := e.GetFlag(name...)
	if f == nil {
		return wv.EMPTY
	}

	return f.GetAsString()
}

// GetAsStringOrRaw will give you the string value of the flag
// with the specified name.
// if there is no flag with this name, or there is an
// error on our path, it will return you the raw data in the
// EventArgs struct.
func (e *EventArgs) GetAsStringOrRaw(name ...string) string {
	tmp := e.GetAsString(name...)

	if wotoStrings.IsEmpty(&tmp) {
		return e.rawData
	}

	return tmp
}

// GetAsStringT's functionality is exactly like
// GetAsString, but it will also trim the spaces,
// so you can use a pure string.
// please consider that using this method is
// somehow dangerous and the usage is for places
// like toLang in wotoTranslation package.
func (e *EventArgs) GetAsStringT(name ...string) string {
	s := e.GetAsString(name...)
	return strings.TrimSpace(s)
}

// GetAsStringTOrRaw's functionality is exactly like
// GetAsString, but it will also trim the spaces,
// so you can use a pure string.
// please consider that using this method is
// somehow dangerous and the usage is for places
// like toLang in wotoTranslation package.
// if the value is empty, it will return the raw data.
func (e *EventArgs) GetAsStringTOrRaw(name ...string) string {
	tmp := e.GetAsStringT(name...)
	if wotoStrings.IsEmpty(&tmp) {
		return e.rawData
	}

	return tmp
}

// GetAsInteger will give you the integer value of the flag
// with the specified name.
// if there is no flag with this name, or there is an
// error on our path, it will return you zero and false.
func (e *EventArgs) GetAsInteger(name ...string) (vI int64, ok bool) {
	f := e.GetFlag(name...)
	if f == nil {
		return wv.BaseIndex, false
	}

	return f.GetAsInteger()
}

// GetAsBool will give you the boolean value of the flag
// with the specified name.
// if there is no flag with this name, or there is an
// error on our path, it will return you false.
func (e *EventArgs) GetAsBool(name ...string) bool {
	f := e.GetFlag(name...)
	if f == nil {
		return false
	}

	return f.GetAsBool()
}

//---------------------------------------------------------

func ToBoolType(value string) (v, isBool bool) {
	value = strings.Trim(value, wv.SPACE_VALUE)
	value = strings.ToLower(value)
	switch value {
	case TrueHlc, YesHlc, OnHlc:
		return true, true
	case FalseHlc, NoHlc, OffHlc:
		return false, true
	default:
		return false, false
	}
}

func (t *FlagType) ToString() string {
	switch *t {
	case NoneFlagType:
		return NoneTypeStr
	case BoolFlagType:
		return BoolTypeStr
	case StringFlagType:
		return StringTypeStr
	case UInt8FlagType:
		return UInt8TypeStr
	case UInt16FlagType:
		return UInt16TypeStr
	case UInt32FlagType:
		return UInt32TypeStr
	case UInt64FlagType:
		return UInt64TypeStr
	case Int8FlagType:
		return Int8TypeStr
	case Int16FlagType:
		return Int16TypeStr
	case Int32FlagType:
		return Int32TypeStr
	case Int64FlagType:
		return Int64TypeStr
	}

	return NoneTypeStr
}

func (t *FlagType) Compare(value *FlagType) bool {
	return false
}

func (t *FlagType) IsInteger() bool {
	return *t == Int8FlagType ||
		*t == Int16FlagType ||
		*t == Int32FlagType ||
		*t == Int64FlagType ||
		*t == UInt8FlagType ||
		*t == UInt16FlagType ||
		*t == UInt32FlagType ||
		*t == UInt64FlagType
}

//---------------------------------------------------------
