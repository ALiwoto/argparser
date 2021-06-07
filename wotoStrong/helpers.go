// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrong

import (
	wv "github.com/ALiwoto/argparser/wotoValues"
)

func repairString(value *string) *string {
	entered := false
	ignoreNext := false
	final := wv.EMPTY
	last := len(*value) - wv.BaseIndex
	next := wv.BaseIndex
	for i, current := range *value {
		if ignoreNext {
			ignoreNext = false
			continue
		}

		if current == wv.CHAR_STR {
			if !entered {
				entered = true
			} else {
				entered = false
			}

			final += string(current)
			continue
		} else {
			if !entered {
				final += string(current)
				continue
			}

			if isSpecial(current) {
				final += wv.BackSlash + string(current)
				continue
			} else {
				if current == wv.LineChar {
					if i != last {
						next = i + wv.BaseOneIndex
						if (*value)[next] == wv.LineChar {
							final += wv.BackSlash +
								string(current) + string(current)
							ignoreNext = true
							continue
						}
					}
				}
			}
		}

		final += string(current)
	}

	return &final
}

func isSpecial(r rune) bool {
	switch r {
	case wv.EqualChar, wv.DPointChar:
		return true
	default:
		return false
	}

}
