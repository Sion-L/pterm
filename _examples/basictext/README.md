### basictext/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/basictext/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/Sion-L/pterm"

func main() {
	// A BasicText printer is used to print text, without special formatting.
	// As it implements the TextPrinter interface, you can use it in combination with other printers.
	pterm.DefaultBasicText.Println("Default basic text printer.")
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")
	pterm.DefaultBasicText.Println("For example to resolve progressbars and spinners.")
	// If you just want to print text, you should use this instead:
	// 	pterm.Println("Hello, World!")
}

```

</details>
