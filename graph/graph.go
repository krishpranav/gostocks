package graph

const (
	dateFormat = "Mon 01/01/2007 15:01"
	timeFormat = "3.01pm"
	dayFormat = "1 Jan"
)

type ChartTheme int

const (
	LineTheme ChartTheme = iota
	DotTheme
	IconTheme
)

func borderHorizontal(out *string, width int) {
	for _i := 0; _i < width-2; _i++ {
		*out += "-"
	}
}