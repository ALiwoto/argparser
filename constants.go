package argparser

const (
	packageErrorSign = "error in argparser package: "
	argFileErrorSign = "arg.go: "
	parseErrMessage  = packageErrorSign + argFileErrorSign +
		"couldn't parse the flags and values.\n stoped in index: "
)
