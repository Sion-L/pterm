# tree/demo

![Animation](animation.svg)

```go
package main

import (
	"github.com/Sion-L/pterm"
)

func main() {
	tree := pterm.TreeNode{
		Text: "Top node",
		Children: []pterm.TreeNode{{
			Text: "Child node",
			Children: []pterm.TreeNode{
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
			},
		}},
	}

	pterm.DefaultTree.WithRoot(tree).Render()
}

```
