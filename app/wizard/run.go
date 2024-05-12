package wizard

import "os"

func Run() {
	if err := root_cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
