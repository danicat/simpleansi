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

// WithBlueBackground wraps the given text with blue background and reset escape sequences.
func WithBlueBackground(text string) string {
	return "\x1b[44m" + text + "\x1b[0m"
}
