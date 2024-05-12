package wizard

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/fatih/color"
	"github.com/mitchellh/go-ps"
)

var (
	itemRequiredColor = color.FgYellow
	futureItemColor   = color.FgMagenta
	itemCreationColor = color.FgGreen
	enterColor        = color.FgCyan
)

func (wiz *Wizard) Warn(args ...interface{}) {
	color.Set(color.FgYellow)
	fmt.Print(wiz.logPrefix)
	fmt.Println(args...)
	color.Unset()
}

func (wiz *Wizard) Ok(args ...interface{}) {
	fmt.Print(wiz.logPrefix)

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

func (wiz *Wizard) consoleHasUnicode() bool {
	switch runtime.GOOS {
	case "windows":
		ppid := os.Getppid()
		proc, err := ps.FindProcess(ppid)
		if err != nil {
			return false
		}

		if strings.HasSuffix(proc.Executable(), "explorer.exe") {
			return false
		}

		if strings.HasSuffix(proc.Executable(), "cmd.exe") {
			return false
		}

		return true
	default:
		return true
	}
}

func (wiz *Wizard) initIO() {
	if wiz.consoleHasUnicode() {
		wiz.logPrefix = "🧙🏿"
	} else {
		wiz.logPrefix = "[Wiz] "
	}
}

func (wiz *Wizard) Fatal(err error) WizFunc {
	return func(w *Wizard, prev WizFunc) WizFunc {
		if errors.Is(err, terminal.InterruptErr) {
			w.Warn("^C again to quit.")
		} else {
			if err != nil {
				w.Warn("A fatal error has occured", err)
			}
		}

		var opt int
		prompt := &survey.Select{
			Message: "What now?",
			Options: []string{"Quit the wizard", "Take me back to where I was.", "Take me back to the main menu."},
		}

		if err := survey.AskOne(prompt, &opt); err == terminal.InterruptErr {
			os.Exit(0)
		}

		switch opt {
		case 0:
			return nil
		case 1:
			return prev
		case 2:
			return SplashScreen
		}

		return nil
	}
}
