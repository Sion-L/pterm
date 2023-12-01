package main

import "github.com/Sion-L/pterm"

func main() {
	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Enter your password")

	logger := pterm.DefaultLogger
	logger.Info("Password received", logger.Args("password", result))
}
