package gobular

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strings"

	rw "github.com/mattn/go-runewidth"
)

type Setup struct {
	GridVBar        rune
	GridHBar        rune
	GridCross       rune
	GridTop         rune
	GridBottom      rune
	GridLeft        rune
	GridRight       rune
	GridTopLeft     rune
	GridTopRight    rune
	GridBottomLeft  rune
	GridBottomRight rune
}

type Table struct {
	Setup
	Caption      string
	EmbedCaption bool
	Rows         []Row
}

type Row struct {
	Cells []Cell
}

type Cell struct {
	// Content is an array of lines in this cell.
	Content []string

	// ColSpan specifies how many cells this one should span. If you have a
	// table with four columns and the first cell should be 2 columns wide, set
	// this to 2 in the first cell and do not use the second cell, it will be
	// ignored. So your Cells array in the row still has four elements. Setting
	// ColSpan to 0 or 1 has the same default effect of just spanning one cell.
	//
	// The WidthMin and WidthMax properties of cells with ColSpan > 1 will be
	// ignoered. Setting restrictions for such cells is not possible currently.
	ColSpan uint
	HAlign  Alignment

	WidthMin uint32
	WidthMax uint32

	// ignore marks a cell ignored due to a ColSpan value from a previous cell.
	ignore bool
}

type Alignment int

const (
	HAlignLeft Alignment = iota
	HAlignCenter
	HAlignRight
)

var (
	DefaultSetup = Setup{'│', '─', '┼', '┬', '┴', '├', '┤', '┌', '┐', '└', '┘'}
	ansiEsc      = regexp.MustCompile("\033.*?m")
)

func NewTable() *Table {
	return &Table{
		Setup: DefaultSetup,
	}
}

func (s *Setup) check(embedCaption bool) bool {
	return rw.StringWidth(string([]rune{
		s.GridVBar,
		s.GridCross,
		s.GridTop,
		s.GridBottom,
		s.GridLeft,
		s.GridRight,
		s.GridTopLeft,
		s.GridTopRight,
		s.GridBottomLeft,
		s.GridBottomRight,
	}))%10 == 0 &&
		(!embedCaption || rw.RuneWidth(s.GridHBar) == rw.RuneWidth(s.GridVBar))
}

func outLen(s string) uint32 {
	return uint32(rw.StringWidth(ansiEsc.ReplaceAllLiteralString(s, "")))
}

func (t *Table) Render() (lines []string, err error) {
	var buf bytes.Buffer
	var colWidths []uint32
	var minWidths []uint32
	var maxWidths []uint32

	if !t.Setup.check(t.EmbedCaption) {
		return nil, fmt.Errorf("Unable to render table: Border and grid runes do not have the same width")
	}

	// calculate column restrictions
	for _, row := range t.Rows {
		for i, cell := range row.Cells {
			// ignore cells which are overshadowed by previous cells with
			// ColSpan > 1
			if cell.ignore {
				continue
			}
			if cell.ColSpan > 1 {
				cell.WidthMin = 0
				cell.WidthMax = 0
				for _, c := range row.Cells[i+1 : i+int(cell.ColSpan)] {
					c.ignore = true
				}
			}

			// find maximum content length in this cell
			maxContentWidth := uint32(0)
			for _, line := range cell.Content {
				if linelen := outLen(line); linelen > maxContentWidth {
					maxContentWidth = linelen
				}
			}

			if len(colWidths) <= i { // grow arrays if we see a new column
				colWidths = append(colWidths, maxContentWidth)
				minWidths = append(minWidths, cell.WidthMin)
				max := uint32(math.MaxUint32)
				if cell.WidthMax != 0 {
					max = cell.WidthMax
				}
				maxWidths = append(maxWidths, max)
			} else { // update limits otherwise
				if maxContentWidth > colWidths[i] {
					colWidths[i] = maxContentWidth
				}
				if cell.WidthMin > minWidths[i] {
					minWidths[i] = cell.WidthMin
				}
				if cell.WidthMax != 0 && cell.WidthMax < maxWidths[i] {
					maxWidths[i] = cell.WidthMax
				}
			}
		}
	}

	tablewidth := uint32(rw.RuneWidth(t.GridVBar))
	for i, w := range colWidths {
		// calculate final column widths
		if minWidths[i] > maxWidths[i] {
			return nil, fmt.Errorf("Unable to render table: WidthMax < WidthMin for column %d (zero based)", i)
		}
		if w > maxWidths[i] {
			w = maxWidths[i]
		}
		if w < minWidths[i] {
			w = minWidths[i]
		}

		// east asian width ambiguous runes (e.g. default boxdrawing) need our
		// cells to be a multiple of 2 characters wide
		if rw.RuneWidth(t.GridVBar) == 2 && w%2 != 0 {
			if w < maxWidths[i] {
				w++
			} else if w > minWidths[i] {
				w--
			} else {
				return nil, fmt.Errorf("Unable to render table: Column %d (zero based) has odd width enforced, but we need an even width due to grid character widths", i)
			}
		}

		tablewidth += w + uint32(rw.RuneWidth(t.GridVBar))
	}

	if 0 == len(t.Rows) || 0 == len(colWidths) || 0 != len(colWidths)%2 {
		t.EmbedCaption = false
	}
	capLen := uint32(outLen(t.Caption))
	capWidth := uint32(capLen + 2*uint32(rw.RuneWidth(t.GridVBar)))
	if 0 != capLen%2 {
		capLen++
		capWidth++
	}
	if t.Caption != "" {
	} //TODO: continue here

	buf.WriteString("bla")
	lines = append(lines, buf.String())
	return
}

func fitPad2(mustLen uint32, align Alignment, format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	realLen := outLen(s)
	delta := int(mustLen) - int(realLen)
	if delta > 0 {
		// need padding
		if HAlignRight == align {
			return strings.Repeat(" ", int(delta)) + s
		} else if HAlignCenter == align {
			ret := strings.Repeat(" ", delta/2) + s
			return ret + strings.Repeat(" ", int(math.Ceil(float64(delta)/2)))
		}

		// Left alignment is the default
		return s + strings.Repeat(" ", int(delta))
	} else if delta < 0 {
		// need trimming
		if HAlignRight == align {
			// cut left hand side
			toks := ansiEsc.Split(s, -1)
			lastTok := toks[len(toks)-1]
			tokLen := outLen(lastTok)

			// extract last escape sequence which will be appended in front of
			// the trimmed return string
			esc := ""
			if escs := ansiEsc.FindAllString(s, -1); escs != nil {
				esc = escs[len(escs)-1]
			}

			if tokLen >= mustLen {
				// last token needs to be trimmed. cut one rune, check, repeat.
				for i := range lastTok {
					newTry := lastTok[i:]
					delta2 := int(mustLen) - int(outLen(newTry))
					if 0 == delta2 {
						return esc + newTry
					} else if 0 < delta2 {
						return esc + strings.Repeat(" ", delta2) + newTry
					}
				}
			} else {
				// last token should be contained completely in return string.
				// recurse to find the token we need to cut.
				escLocs := ansiEsc.FindAllStringIndex(s, -1)
				loc := escLocs[len(escLocs)-1][0]
				return fitPad2(mustLen-tokLen, align, s[0:loc]) + esc + lastTok
			}
		} else if HAlignCenter == align {
			// first cut left, then cut right
			rhs := fitPad2(mustLen+uint32(-delta/2), HAlignRight, s)
			return fitPad2(mustLen, HAlignLeft, rhs)
		} else {
			// Left alignment is the default
			// cut right hand side
			toks := ansiEsc.Split(s, 2)
			tokLen := outLen(toks[0])
			if tokLen >= mustLen {
				// first token must be trimmed
				return fmt.Sprintf("%.*s", mustLen, toks[0])
			}

			// first token should be contained completely in return string.
			// recurse to find the token we need to cut.
			esc := ansiEsc.FindString(s)
			return toks[0] + esc + fitPad2(mustLen-tokLen, align, toks[1])
		}
	}

	// perfect length
	return s
}

func fitPad(mustLen uint32, align Alignment, format string, a ...interface{}) string {
	return fitPad2(mustLen, align, format, a...) + "\033[0m"
}
