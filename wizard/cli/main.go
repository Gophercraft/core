package cli

import "os"

func Run() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
