package main

import "github.com/Sion-L/pterm"

func main() {
	pterm.DefaultBox.
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(2).
		WithBottomPadding(2).
		Println("Hello, World!")
}
