package golib

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var policy *bluemonday.Policy

func RenderMD(md string) string {
	unsafe := blackfriday.Run([]byte(md))
	html := policy.SanitizeBytes(unsafe)
	return string(html)
}

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateGUID() string {
	return uuid.NewV4().String()
}

func init() {
	rand.Seed(time.Now().UnixNano())
	policy = bluemonday.UGCPolicy()
}

type TableOpts struct {
	Title         string
	Headers       []string
	Padding       int
	DrawRowBorder bool
}

func DefaultTableOpts() TableOpts {
	return TableOpts{
		Title:         "",
		Headers:       []string{},
		Padding:       1,
		DrawRowBorder: false,
	}
}

func ToTable(rows [][]string, opts TableOpts) string {
	colSize := map[int]int{}
	for r := range rows {
		for c, col := range rows[r] {
			if len(strings.TrimSpace(col)) > colSize[c] {
				colSize[c] = len(strings.TrimSpace(col))
			}
		}
	}
	totalSize := 0
	// column data size + padding for each side
	for i := range colSize {
		totalSize += colSize[i] + (2 * opts.Padding)
	}
	// spacers
	totalSize += len(colSize) - 1
	// sides
	totalSize += 2
	fullRowLine := "+" + strings.Repeat("-", totalSize-2) + "+\n"
	m := fullRowLine
	if len(opts.Title) > 0 {
		p := totalSize/2 - len(strings.TrimSpace(opts.Title))/2
		formatStr := fmt.Sprintf("|%s%%-%dv%s|\n",
			strings.Repeat(" ", p-1),
			len(strings.TrimSpace(opts.Title)),
			strings.Repeat(" ", p-1),
		)
		m += fmt.Sprintf(formatStr, strings.TrimSpace(opts.Title))
		m += fullRowLine
	}
	for r := range rows {
		rowStr := "|"
		for c, col := range rows[r] {
			if c > 0 {
				rowStr += "|"
			}
			formatStr := fmt.Sprintf("%s%%-%dv%s",
				strings.Repeat(" ", opts.Padding),
				colSize[c],
				strings.Repeat(" ", opts.Padding),
			)
			rowStr += fmt.Sprintf(formatStr, strings.TrimSpace(col))
		}
		rowStr += "|\n"
		if opts.DrawRowBorder {
			rowStr += fullRowLine
		}
		m += rowStr

	}
	if !opts.DrawRowBorder {
		m += fullRowLine
	}
	return m
}
