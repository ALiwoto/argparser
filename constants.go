// argparser Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package argparser

const (
	packageErrorSign = "error in argparser package: "
	argFileErrorSign = "arg.go: "
	parseErrMessage  = packageErrorSign + argFileErrorSign +
		"couldn't parse the flags and values.\n stoped in index: "
)
