// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoValues

import "time"

// router config values
const (
	APP_PORT        = "PORT"
	GET_SLASH       = "/"
	HTTP_ADDRESS    = ":"
	FORMAT_VALUE    = "%v"
	SPACE_VALUE     = " "
	LineEscape      = "\n"
	R_ESCAPE        = "\r"
	SEMICOLON       = ";"
	CAMA            = ","
	ParaOpen        = "("
	ParaClose       = ")"
	NullStr         = "null"
	DoubleQ         = "\""
	SingleQ         = "'"
	DoubleQJ        = "”"
	BracketOpen     = "["
	Bracketclose    = "]"
	Star            = "*"
	BackSlash       = "\\"
	DoubleBackSlash = "\\\\"
	Point           = "."
	AutoStr         = "auto"
	AtSign          = "@"
	EqualStr        = "="
	DdotSign        = ":"
)

const (
	AcceptKey     = "Accept"
	AcceptAll     = "*/*"
	EnUs          = "en-US"
	ConnectionKey = "Connection"
	KeepAlive     = "keep-alive"
	RefererKey    = "Referer"
	Get           = "GET"
)

const TICK_INTERVAL = 25 * time.Second // seconds
