// Bot.go Project
// Copyright (C) 2021 Sayan Biswas, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package argparser

import "strings"

// High-level constant.
type Hlc struct {
	Name  string
	Value interface{}
	Type  FlagType
}

type HlcCollection []Hlc

//---------------------------------------------------------

// common and global hlcs.
const (
	TrueHlc  = "true"
	YesHlc   = "yes"
	OnHlc    = "on"
	FalseHlc = "false"
	NoHlc    = "no"
	OffHlc   = "off"
)

//---------------------------------------------------------

func GetHlcCollection(h ...Hlc) HlcCollection {
	return h
}

func getInternal(name string) (interface{}, FlagType) {
	// about this I have to tell you something.
	// internal Hlcs are allowed to be case insensetive.
	// for example if a user type
	// --del=TRUE
	// it's equal to:
	// --del=true
	// also see these forms:
	// --del=ON, --del=on, --del:on, --del on
	// --del=YES, --del=yes, --del:yEs, --del yeS
	// --del=NO, --del=no, --del:nO, --del No
	// --del=OFF, --del=off, --del:ofF, --del Off
	// for boolean values, if user doesn't write
	// `true` or `false`, it will be considered as
	// `true`. for example:
	// `/pat --del`
	// is equal to:
	// `/pat --del=true` or `/pat --del=on`,
	// it's not that if the default value of
	// a flag is `false`, then it will become
	// oppisite and it will become `true`.
	// even if the default value of a param is
	// `true`, the define won't make it `false`,
	// means it doesn't have any effect to use a
	// flag param of boolean in a command.
	switch strings.ToLower(name) {
	case TrueHlc, YesHlc, OnHlc:
		return true, BoolFlagType
	case FalseHlc, NoHlc, OffHlc:
		return false, BoolFlagType
	default:
		return nil, NoneFlagType
	}
}

//---------------------------------------------------------

func (c *HlcCollection) Contains(name string) bool {
	for _, h := range *c {
		if h.Name == name {
			return true
		}
	}

	return false
}

func (c *HlcCollection) GetValue(name string) (interface{}, FlagType) {
	for _, h := range *c {
		if h.Name == name {
			return h.Value, h.Type
		}
	}

	v, t := getInternal(name)
	if t != NoneFlagType {
		return v, t
	}

	return nil, NoneFlagType
}

//---------------------------------------------------------
