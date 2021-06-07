// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

import (
	"reflect"
	"strings"

	tf "github.com/ALiwoto/argparser/interfaces"
	"github.com/ALiwoto/argparser/wotoStrings"
	wv "github.com/ALiwoto/argparser/wotoValues"
)

// GetValue will give you the real value of this StrongString.
func (_s *StrongString) GetValue() string {
	return string(_s._value)
}

// length method, will give you the length-as-int of this StrongString.
func (_s *StrongString) Length() int {
	return len(_s._value)
}

// isEmpty will check if this StrongString is empty or not.
func (_s *StrongString) IsEmpty() bool {
	return _s._value == nil || len(_s._value) == wv.BaseIndex
}

// isEqual will check if the passed-by-value in the arg is equal to this
// StrongString or not.
func (_s *StrongString) IsEqual(_q tf.QString) bool {
	if reflect.TypeOf(_q) != reflect.TypeOf(_s) {
		return _q.GetValue() == _s.GetValue()
	}

	_strong, _ok := _q.(*StrongString)
	if !_ok {
		return false
	}
	// check if the length of them are equal or not.
	if len(_s._value) != len(_strong._value) {
		//fmt.Println(len(_s._value), len(_strong._value))
		return false
	}
	for i := 0; i < len(_s._value); i++ {
		if _s._value[i] != _strong._value[i] {
			//fmt.Println(_s._value[i], _strong._value[i])
			return false
		}
	}
	return true
}

// GetIndexV method will give you the rune in _index.
func (_s *StrongString) GetIndexV(_index int) rune {
	if _s.IsEmpty() {
		return wv.BaseIndex
	}

	l := len(_s._value)

	if _index >= l || l < wv.BaseIndex {

		return _s._value[wv.BaseIndex]
	}

	return _s._value[_index]
}

// HasSuffix will check if at least there is one suffix is
// presents in this StrongString not.
// the StrongString should ends with at least one of these suffixes.
func (_s *StrongString) HasSuffix(values ...string) bool {
	for _, s := range values {
		if strings.HasSuffix(_s.GetValue(), s) {
			return true
		}
	}

	return false
}

// HasSuffixes will check if all of the suffixes are
// present in this StrongString or not.
// the StrongString should ends with all of these suffixes.
// usage of this method is not recommended, since you can use
// HasSuffix method with only one string (the longest string).
// this way you will just use too much cpu resources.
func (_s *StrongString) HasSuffixes(values ...string) bool {
	for _, s := range values {
		if !strings.HasSuffix(_s.GetValue(), s) {
			return false
		}
	}

	return true
}

// HasPrefix will check if at least there is one prefix is
// presents in this StrongString or not.
// the StrongString should starts with at least one of these prefixes.
func (_s *StrongString) HasPrefix(values ...string) bool {
	for _, s := range values {
		if strings.HasPrefix(_s.GetValue(), s) {
			return true
		}
	}

	return false
}

// HasPrefixes will check if all of the prefixes are
// present in this StrongString or not.
// the StrongString should starts with all of these suffixes.
// usage of this method is not recommended, since you can use
// HasSuffix method with only one string (the longest string).
// this way you will just use too much cpu resources.
func (_s *StrongString) HasPrefixes(values ...string) bool {
	for _, s := range values {
		if !strings.HasPrefix(_s.GetValue(), s) {
			return false
		}
	}

	return true
}

func (_s *StrongString) Split(qs ...tf.QString) []tf.QString {
	strs := wotoStrings.SplitSlice(_s.GetValue(), ToStrSlice(qs))
	return ToQSlice(strs)
}

func (_s *StrongString) SplitN(n int, qs ...tf.QString) []tf.QString {
	strs := wotoStrings.SplitSliceN(_s.GetValue(), ToStrSlice(qs), n)
	return ToQSlice(strs)
}

func (_s *StrongString) SplitFirst(qs ...tf.QString) []tf.QString {
	strs := wotoStrings.SplitSliceN(_s.GetValue(), ToStrSlice(qs), wv.BaseTwoIndex)
	return ToQSlice(strs)
}

func (_s *StrongString) SplitStr(qs ...string) []tf.QString {
	strs := wotoStrings.SplitSlice(_s.GetValue(), qs)
	return ToQSlice(strs)
}

func (_s *StrongString) SplitStrN(n int, qs ...string) []tf.QString {
	strs := wotoStrings.SplitSliceN(_s.GetValue(), qs, n)
	return ToQSlice(strs)
}

func (_s *StrongString) SplitStrFirst(qs ...string) []tf.QString {
	strs := wotoStrings.SplitSliceN(_s.GetValue(), qs, wv.BaseTwoIndex)
	return ToQSlice(strs)
}

func (_s *StrongString) ToQString() tf.QString {
	return _s
}

func (_s *StrongString) Contains(qs ...tf.QString) bool {
	v := _s.GetValue()
	for _, current := range qs {
		if strings.Contains(v, current.GetValue()) {
			return true
		}
	}

	return false
}

func (_s *StrongString) ContainsStr(str ...string) bool {
	v := _s.GetValue()
	for _, current := range str {
		if strings.Contains(v, current) {
			return true
		}
	}

	return false
}

func (_s *StrongString) ContainsAll(qs ...tf.QString) bool {
	v := _s.GetValue()
	for _, current := range qs {
		if !strings.Contains(v, current.GetValue()) {
			return false
		}
	}

	return true
}

func (_s *StrongString) ContainsStrAll(str ...string) bool {
	v := _s.GetValue()
	for _, current := range str {
		if !strings.Contains(v, current) {
			return false
		}
	}

	return true
}

func (_s *StrongString) TrimPrefix(qs ...tf.QString) tf.QString {
	final := _s.GetValue()
	for _, current := range qs {
		final = strings.TrimPrefix(final, current.GetValue())
	}

	return SsPtr(final)
}

func (_s *StrongString) TrimPrefixStr(qs ...string) tf.QString {
	final := _s.GetValue()
	for _, current := range qs {
		final = strings.TrimPrefix(final, current)
	}

	return SsPtr(final)
}

func (_s *StrongString) TrimSuffix(qs ...tf.QString) tf.QString {
	final := _s.GetValue()
	for _, current := range qs {
		final = strings.TrimSuffix(final, current.GetValue())
	}

	return SsPtr(final)
}

func (_s *StrongString) TrimSuffixStr(qs ...string) tf.QString {
	final := _s.GetValue()
	for _, current := range qs {
		final = strings.TrimSuffix(final, current)
	}

	return SsPtr(final)
}

func (_s *StrongString) Trim(qs ...tf.QString) tf.QString {
	final := _s.GetValue()
	for _, current := range qs {
		final = strings.Trim(final, current.GetValue())
	}

	return SsPtr(final)
}

func (_s *StrongString) TrimStr(qs ...string) tf.QString {
	final := _s.GetValue()
	for _, current := range qs {
		final = strings.Trim(final, current)
	}

	return SsPtr(final)
}

func (_s *StrongString) Replace(qs, newS tf.QString) tf.QString {
	return _s.ReplaceStr(qs.GetValue(), newS.GetValue())
}

func (_s *StrongString) ReplaceStr(qs, newS string) tf.QString {
	final := _s.GetValue()
	final = strings.ReplaceAll(final, qs, newS)
	return SsPtr(final)
}

// LockSpecial will lock all the defiend special characters.
// This way, you don't actually have to be worry about
// some normal mistakes in spliting strings, cut them out,
// check them. join them, etc...
// WARNING: this method is so dangerous, it's really
// dangerous. we can't say that it's unsafe actually,
// but still it's really dangerous, so if you don't know what the
// fuck are you doing, then please don't use this method.
// this method will not return you a new value, it will effect the
// current value. please consider using it carefully.
func (_s *StrongString) LockSpecial() {
	final := _s.GetValue()
	// replacing escaped string characters
	// (I mean escaped double quetion mark) is necessary before
	// repairing value.
	final = strings.ReplaceAll(final, wv.BACK_STR, wv.JA_STR)

	// let it repair the string.
	// this function is for repairing these special characters
	// and strings:
	// '=', ':' and "=="
	// it will escape them.
	// if it wasn't for this function, members had to
	// escape all of these bullshits themselves...
	// hahaha, you see, it's actually usefull.
	final = *repairString(&final)

	final = strings.ReplaceAll(final, wv.BACK_FLAG, wv.JA_FLAG)
	final = strings.ReplaceAll(final, wv.BACK_EQUALITY, wv.JA_EQUALITY)
	final = strings.ReplaceAll(final, wv.BACK_DDOT, wv.JA_DDOT)

	_s._value = make([]rune, wv.BaseIndex)
	for _, c := range final {
		if c != wv.BaseIndex {
			_s._value = append(_s._value, c)
		}
	}
}

// UnlockSpecial will unlock all the defiend special characters.
// it will return them to their normal form.
func (_s *StrongString) UnlockSpecial() {
	final := _s.GetValue()
	final = strings.ReplaceAll(final, wv.JA_FLAG, wv.FLAG_PREFIX)
	final = strings.ReplaceAll(final, wv.JA_STR, wv.STR_SIGN)
	final = strings.ReplaceAll(final, wv.JA_EQUALITY, wv.EqualStr)
	final = strings.ReplaceAll(final, wv.JA_DDOT, wv.DdotSign)

	_s._value = make([]rune, wv.BaseIndex)
	for _, c := range final {
		if c != wv.BaseIndex {
			_s._value = append(_s._value, c)
		}
	}
}
