// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoStrings

import (
	"strings"

	wv "github.com/ALiwoto/argparser/wotoValues"
)

func Split(_s string, separator ...string) []string {
	if len(separator) == wv.BaseIndex {
		return nil
	}

	final := _s
	for _, myStr := range separator {
		final = strings.ReplaceAll(final, myStr, sepStr)
	}
	return FixSplit(strings.Split(final, sepStr))
}

func SplitN(_s string, n int, separator ...string) []string {
	if len(separator) == wv.BaseIndex {
		return nil
	}

	rep := n - wv.BaseOneIndex
	final := _s
	done := wv.BaseIndex

	for _, myStr := range separator {
		if done < rep {
			if strings.Contains(final, myStr) {
				final = strings.Replace(final, myStr, sepStr, rep)
				done++
			}
		} else {
			break
		}

	}

	theS := strings.SplitN(final, sepStr, n)

	return FixSplit(theS)
}

func SplitSlice(_s string, separator []string) []string {
	if len(separator) == wv.BaseIndex {
		return nil
	}

	final := _s
	for _, myStr := range separator {
		final = strings.ReplaceAll(final, myStr, sepStr)
	}

	return FixSplit(strings.Split(final, sepStr))
}

func SplitSliceN(_s string, separator []string, n int) []string {
	if len(separator) == wv.BaseIndex {
		return nil
	}

	rep := n - wv.BaseOneIndex
	final := _s
	done := wv.BaseIndex

	for _, myStr := range separator {
		if done < rep {
			if strings.Contains(final, myStr) {
				final = strings.Replace(final, myStr, sepStr, rep)
				done++
			}
		} else {
			break
		}

	}

	theS := strings.SplitN(final, sepStr, n)

	return FixSplit(theS)
}

// FixSplit will fix the bullshit bug in the
// Split function (which is not ignoring the spaces between strings).
func FixSplit(myStrings []string) []string {
	final := make([]string, wv.BaseIndex, cap(myStrings))

	for _, current := range myStrings {
		if !IsEmpty(&current) {
			final = append(final, current)
		}
	}

	return final
}

// IsEmpty function will check if the passed-by
// string value is empty or not.
func IsEmpty(_s *string) bool {
	return len(*_s) == wv.BaseIndex
}

// YesOrNo returns yes if v is true, otherwise no,
func YesOrNo(v bool) string {
	if v {
		return wv.Yes
	} else {
		return wv.No
	}
}
