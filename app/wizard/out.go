package wizard

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	item_required_color = color.FgYellow
	future_item_color   = color.FgMagenta
	item_creation_color = color.FgGreen
	enterColor          = color.FgCyan
)

func printwarn(args ...any) {
	color.Set(color.FgRed)
	fmt.Print("🧙🏿")
	fmt.Println(args...)
	color.Unset()
}

func printfok(f string, args ...any) {
	printok(fmt.Sprintf(f, args...))
}

func printok(args ...any) {
	fmt.Print("🧙🏿")

	for i, v := range args {
		switch attr := v.(type) {
		case color.Attribute:
			color.Set(attr)
		default:
			if i > 0 {
				fmt.Print(" ")
			}

			fmt.Print(v)
		}
	}

	fmt.Println()
	color.Unset()
}
