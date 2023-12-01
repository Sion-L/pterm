# bigtext/colored

![Animation](animation.svg)

```go
package main

import (
	"github.com/Sion-L/pterm"
	"github.com/Sion-L/pterm/putils"
)

func main() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle())).
		Render()
}

```
