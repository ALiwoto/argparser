// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

import wv "github.com/ALiwoto/argparser/wotoValues"

// _setValue will set the bytes value of the StrongString.
func (_s *StrongString) _setValue(str string) {
	if _s._value == nil {
		_s._value = make([]rune, wv.BaseIndex)
	}

	for _, current := range str {
		_s._value = append(_s._value, current)
	}
}
