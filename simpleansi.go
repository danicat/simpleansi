package simpleansi

import "fmt"

// ClearScreen cleans the terminal and set cursor position to the top left corner.
func ClearScreen() {
	fmt.Print("\x1b[2J")
	MoveCursor(0, 0)
}

// MoveCursor sets the cursor position to given row and col.
//
// Please note that ANSI is 1-based and the top left corner is (1,1), but here we are assuming
// the user is using a zero based coordinate system where the top left corner is (0, 0)
func MoveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}

const reset = "\x1b[0m"

type Colour int

const (
	BLACK Colour = iota
	RED
	GREEN
	BROWN
	BLUE
	MAGENTA
	CYAN
	GREY
)

var colours = map[Colour]string{
	BLACK:   "\x1b[1;30;40m",
	RED:     "\x1b[1;31;41m",
	GREEN:   "\x1b[1;32;42m",
	BROWN:   "\x1b[1;33;43m",
	BLUE:    "\x1b[1;34;44m",
	MAGENTA: "\x1b[1;35;45m",
	CYAN:    "\x1b[1;36;46m",
	GREY:    "\x1b[1;37;47m",
}

// WithBlueBackground wraps the given text with blue background and reset escape sequences.
func WithBlueBackground(text string) string {
	return "\x1b[44m" + text + reset
}

// WithBackground wraps the given 'text' with 'colour' background and reset escape sequences.
func WithBackground(text string, colour Colour) string {
	if c, ok := colours[colour]; ok {
		return c + text + reset
	}
	//Default to blue if none resolved
	return WithBlueBackground(text)
}
