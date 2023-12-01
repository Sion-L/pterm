# bigtext/default

![Animation](animation.svg)

```go
package main

import (
	"github.com/Sion-L/pterm"
	"github.com/Sion-L/pterm/putils"
)

func main() {
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()
}

```
