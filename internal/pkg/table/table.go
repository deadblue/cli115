/*
A simple replacement for "github.com/olekukonko/tablewriter"
*/
package table

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"io"
	"strings"
)

type ColumnAlign int

const (
	AlignLeft ColumnAlign = iota
	AlignRight

	charSpace  = " "
	charCross  = "+"
	charRowSep = "-"
	charColSep = "|"
)

type column struct {
	name  string
	aligh ColumnAlign
	width int
}

type Table struct {
	cols    []*column
	colSize int
	rows    [][]string
}

/*
Add a column definition.
*/
func (t *Table) AddColumn(name string, align ColumnAlign) *Table {
	if t.cols == nil {
		t.cols = make([]*column, 0)
	}
	t.cols = append(t.cols, &column{
		name:  name,
		aligh: align,
		width: runewidth.StringWidth(name),
	})
	t.colSize += 1
	return t
}

/*
Append datas as a row.
If data count is more than the column count, exceeded ones will be dropped.
*/
func (t *Table) AppendRow(data []string) {
	if t.colSize == 0 {
		return
	}
	dataSize, row := len(data), make([]string, t.colSize)
	for index, col := range t.cols {
		if index < dataSize {
			row[index] = data[index]
			// Update column width
			width := runewidth.StringWidth(data[index])
			if width > col.width {
				col.width = width
			}
		} else {
			row[index] = ""
		}
	}
	t.rows = append(t.rows, row)
}

/*
Render table to writer.
*/
func (t *Table) Render(w io.Writer) {
	headers, lineBuf := make([]string, t.colSize), &strings.Builder{}
	for index, col := range t.cols {
		headers[index] = col.name
		if index == 0 {
			lineBuf.WriteString(charCross)
		}
		lineBuf.WriteString(strings.Repeat(charRowSep, col.width+2))
		lineBuf.WriteString(charCross)
	}
	line := lineBuf.String()

	_, _ = fmt.Fprintln(w, line)
	t.renderRow(w, headers)
	_, _ = fmt.Fprintln(w, line)
	for _, row := range t.rows {
		t.renderRow(w, row)
	}
	_, _ = fmt.Fprintln(w, line)
}

func (t *Table) renderRow(w io.Writer, row []string) {
	cellCount := len(row)
	for index, col := range t.cols {
		cell, width := "", 0
		if index < cellCount {
			cell, width = row[index], runewidth.StringWidth(row[index])
		}
		padLeft, padRight := "", ""
		if col.aligh == AlignLeft {
			padRight = strings.Repeat(charSpace, col.width-width)
		} else if col.aligh == AlignRight {
			padLeft = strings.Repeat(charSpace, col.width-width)
		}

		if index == 0 {
			_, _ = fmt.Fprint(w, charColSep)
		}
		_, _ = fmt.Fprint(w, charSpace, padLeft, cell, padRight, charSpace, charColSep)
	}
	_, _ = fmt.Fprint(w, "\n")
}

/*
Format table to string.
*/
func (t *Table) String() string {
	buf := &strings.Builder{}
	t.Render(buf)
	return buf.String()
}

/*
Create table.
*/
func New() *Table {
	return &Table{}
}
