package template

// Heavily copied from https://github.com/manifoldco/promptui

import (
	"fmt"
	"strconv"
	"strings"
)

const esc = "\033["

type attribute int

// The possible state of text inside the application, either Bold, faint, italic or underline.
//
// These constants are called through the use of the Styler function.
const (
	reset attribute = iota

	FGBold
	FGFaint
	FGItalic
	FGUnderline
)

// The possible colors of text inside the application.
//
// These constants are called through the use of the Styler function.
const (
	FGBlack attribute = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
)

// The possible background colors of text inside the application.
//
// These constants are called through the use of the Styler function.
const (
	BGBlack attribute = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
)

// ResetCode is the character code used to reset the terminal formatting
var ResetCode = fmt.Sprintf("%s%dm", esc, reset)

const (
	hideCursor = esc + "?25l"
	showCursor = esc + "?25h"
	clearLine  = esc + "2K"
)

// Styler is a function that accepts multiple possible styling transforms from the state,
// color and background colors constants and transforms them into a templated string
// to apply those styles in the CLI.
//
// The returned styling function accepts a string that will be extended with
// the wrapping function's styling attributes.
func Styler(attrs ...attribute) func(interface{}) string {
	a := make([]string, len(attrs))
	for i, v := range attrs {
		a[i] = strconv.Itoa(int(v))
	}

	seq := strings.Join(a, ";")

	return func(v interface{}) string {
		end := ""
		s, ok := v.(string)
		if !ok || !strings.HasSuffix(s, ResetCode) {
			end = ResetCode
		}
		return fmt.Sprintf("%s%sm%v%s", esc, seq, v, end)
	}
}
