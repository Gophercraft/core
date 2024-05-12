package main

import (
	"os"

	"fmt"

	"github.com/Gophercraft/core/format/bls"
	"github.com/Gophercraft/log"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "read",
	Short: "Read BLS shader file",
	Long:  `Decode a BLS shader file, outputting information contained within`,

	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[1]

		fmt.Println("Opening", filePath)

		file, err := os.ReadFile(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println("read file data")

		shader, err := bls.Read(file)
		if err != nil {
			fmt.Println("Bls error: ", err)
			return
		}

		fmt.Println(spew.Sdump(shader))
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.core.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

func main() {
	rootCmd.Execute()
}
