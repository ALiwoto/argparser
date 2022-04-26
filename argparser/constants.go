// argparser Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package argparser

const (
	NoneFlagType FlagType = iota
	BoolFlagType
	StringFlagType
	UInt8FlagType
	UInt16FlagType
	UInt32FlagType
	UInt64FlagType
	Int8FlagType
	Int16FlagType
	Int32FlagType
	Int64FlagType
)

const (
	NoneTypeStr   = "None"
	BoolTypeStr   = "bool"
	StringTypeStr = "string"
	UInt8TypeStr  = "uint8"
	UInt16TypeStr = "uint16"
	UInt32TypeStr = "uint32"
	UInt64TypeStr = "uint64"
	Int8TypeStr   = "int8"
	Int16TypeStr  = "int16"
	Int32TypeStr  = "int32"
	Int64TypeStr  = "int64"
)

// common and global HLCs.
const (
	TrueHlc  = "true"
	YesHlc   = "yes"
	OnHlc    = "on"
	FalseHlc = "false"
	NoHlc    = "no"
	OffHlc   = "off"
)
