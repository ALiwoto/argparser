// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

import (
	tfc "github.com/ALiwoto/argparser/interfaces"
)

const (
	MaxUIDIndex    = 9
	BaseUIDIndex   = 1
	UIDIndexOffSet = 1
)

const (
	StrongOffSet = 13
)

// the StrongString used in the program for High-security!
type StrongString struct {
	_value []rune
}

// Ss will generate a new StrongString
// with the specified non-encoded string value.
func Ss(_s string) StrongString {
	_strong := StrongString{}
	_strong._setValue(_s)
	return _strong
}

// Qss will generate a new QString
// with the specified non-encoded string value.
func Qss(_s string) tfc.QString {
	str := Ss(_s)
	return &str
}

// Sb will generate a new StrongString
// with the specified non-encoded bytes value.
func Sb(_b []byte) StrongString {
	return Ss(string(_b))
}

// QSb will generate a new QString
// with the specified non-encoded bytes value.
func Qsb(_b []byte) tfc.QString {
	str := Ss(string(_b))
	return &str
}

// SS will generate a new StrongString
// with the specified non-encoded string value.
func SsPtr(_s string) *StrongString {
	_strong := StrongString{}
	_strong._setValue(_s)
	return &_strong
}

func ToStrSlice(qs []tfc.QString) []string {
	tmp := make([]string, len(qs))
	for i, current := range qs {
		tmp[i] = current.GetValue()
	}
	return tmp
}

func ToQSlice(strs []string) []tfc.QString {
	tmp := make([]tfc.QString, len(strs))
	for i, current := range strs {
		tmp[i] = SsPtr(current)
	}
	return tmp
}
