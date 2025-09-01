package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	for _, band := range bands {
		printBand(band)
	}
}

type Band struct {
	MHz      float64
	Length   uint64 // cm
	Start    float64
	End      float64
	Segments []Segment
}

type Segment struct {
	Start float64
	End   float64
	Mode  Mode
}

type Mode int

const (
	ALL Mode = iota
	CW
	CWDX
	CWSSB
	DIGI
	FONE
	SAT
	NAR
	BEAC
)

var modeNames = map[Mode]string{
	ALL:   "ALL",
	CW:    "CW",
	CWDX:  "CWDX",
	CWSSB: "CWSSB",
	DIGI:  "DIGI",
	FONE:  "FONE",
	SAT:   "SAT",
	NAR:   "NAR",
	BEAC:  "BEAC",
}

var modeColors = map[Mode]string{
	ALL:   "\033[32m", // Green
	CW:    "\033[32m", // Green
	CWDX:  "\033[32m", // Green
	CWSSB: "\033[36m", // Teac
	DIGI:  "\033[34m", // Blue
	FONE:  "\033[31m", // Red
	SAT:   "\033[35m", // Magenta
	NAR:   "\033[35m", // Magenta
	BEAC:  "\033[33m", // Yellow
}

const resetColor = "\033[0m"

var bands = []Band{
	{
		MHz:    1.8,
		Length: 16000,
		Start:  1810,
		End:    2000,
		Segments: []Segment{
			{Start: 1810, End: 1838, Mode: CW},
			{Start: 1838, End: 1843, Mode: DIGI},
			{Start: 1840, End: 2000, Mode: FONE},
		},
	},
	{
		MHz:    3.5,
		Length: 8000,
		Start:  3500,
		End:    3800,
		Segments: []Segment{
			{Start: 3500, End: 3510, Mode: CWDX},
			{Start: 3500, End: 3580, Mode: CW},
			{Start: 3580, End: 3620, Mode: DIGI},
			{Start: 3600, End: 3800, Mode: FONE},
		},
	},
	{
		MHz:    7,
		Length: 4000,
		Start:  7000,
		End:    7200,
		Segments: []Segment{
			{Start: 7000, End: 7040, Mode: CW},
			{Start: 7040, End: 7060, Mode: DIGI},
			{Start: 7050, End: 7200, Mode: FONE},
		},
	},
	{
		MHz:    10,
		Length: 3000,
		Start:  10100,
		End:    10150,
		Segments: []Segment{
			{Start: 10140, End: 10150, Mode: DIGI},
		},
	},
	{
		MHz:    14,
		Length: 2000,
		Start:  14000,
		End:    14350,
		Segments: []Segment{
			{Start: 14000, End: 14070, Mode: CW},
			{Start: 14070, End: 14099, Mode: DIGI},
			{Start: 14101, End: 14350, Mode: FONE},
		},
	},
	{
		MHz:    18,
		Length: 1700,
		Start:  18068,
		End:    18168,
		Segments: []Segment{
			{Start: 18068, End: 18095, Mode: CW},
			{Start: 18095, End: 18109, Mode: DIGI},
			{Start: 18111, End: 18168, Mode: FONE},
		},
	},
	{
		MHz:    21,
		Length: 1500,
		Start:  21000,
		End:    21450,
		Segments: []Segment{
			{Start: 21000, End: 21070, Mode: CW},
			{Start: 21070, End: 21120, Mode: DIGI},
			{Start: 21151, End: 21450, Mode: FONE},
		},
	},
	{
		MHz:    24,
		Length: 1200,
		Start:  24890,
		End:    24990,
		Segments: []Segment{
			{Start: 24890, End: 24915, Mode: CW},
			{Start: 24915, End: 24929, Mode: DIGI},
			{Start: 24931, End: 24990, Mode: FONE},
		},
	},
	{
		MHz:    28,
		Length: 1000,
		Start:  28000,
		End:    29700,
		Segments: []Segment{
			{Start: 28000, End: 28070, Mode: CW},
			{Start: 28070, End: 28190, Mode: DIGI},
			{Start: 28225, End: 29300, Mode: FONE},
			{Start: 29300, End: 29510, Mode: SAT},
		},
	},
	{
		MHz:    50,
		Length: 600,
		Start:  50000,
		End:    52000,
		Segments: []Segment{
			{Start: 50000, End: 50100, Mode: CW},
			{Start: 50100, End: 52000, Mode: NAR},
			{Start: 51410, End: 51590, Mode: FONE},
		},
	},
	{
		MHz:    144,
		Length: 200,
		Start:  144000,
		End:    146000,
		Segments: []Segment{
			{Start: 144000, End: 144110, Mode: CW},
			{Start: 145806, End: 146000, Mode: SAT},
			{Start: 144400, End: 144490, Mode: BEAC},
		},
	},
	{
		MHz:    430,
		Length: 70,
		Start:  430000,
		End:    440000,
		Segments: []Segment{
			{Start: 432025, End: 432100, Mode: CW},
			{Start: 432400, End: 432490, Mode: BEAC},
			{Start: 432500, End: 432975, Mode: ALL},
			{Start: 432100, End: 432400, Mode: CWSSB},
		},
	},
}

func printBand(band Band) {
	const width = 120

	length := fmt.Sprintf("%dm", band.Length/100)
	if band.Length < 100 {
		length = fmt.Sprintf("%dcm", band.Length)
	}

	fmt.Printf("Band: %.1f Mhz %s\n", band.MHz, length)

	min_ := band.Start
	max_ := band.End
	for _, s := range band.Segments {
		if s.Start < min_ {
			min_ = s.Start
		}
		if s.End > max_ {
			max_ = s.End
		}
	}

	scale := float64(width) / (max_ - min_)

	// Print the full line
	fmt.Printf("Full %6.0f├%s┤%-6.0f\n", min_, strings.Repeat("─", width+1), max_)

	for _, s := range band.Segments {
		startCol := int(math.Round((s.Start - min_) * scale))
		endCol := int(math.Round((s.End - min_) * scale))

		if endCol <= startCol {
			endCol = startCol + 1
		}

		line := make([]rune, width)
		for i := range line {
			line[i] = ' '
		}

		for i := startCol; i < endCol; i++ {
			line[i] = '─'
		}

		startStr := ""
		if startCol > 0 {
			startStr = fmt.Sprintf("%*d", startCol+7, int(s.Start))
		} else {
			startStr = fmt.Sprintf("%6d", int(s.Start))
		}

		endStr := fmt.Sprintf("%d", int(s.End))

		fmt.Printf("%s%-4s %s├%s┤%s%s\n",
			modeColors[s.Mode],
			modeNames[s.Mode],
			startStr,
			string(line[startCol:endCol]),
			endStr,
			resetColor,
		)
	}

	fmt.Println()
}
