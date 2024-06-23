package wizard

import (
	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var home_backup_cmd = &cobra.Command{
	Use:   "backup",
	Short: "(ADMIN) take a backup of the remote home server database",
	Long:  "Create a zip archive image of the entire remote home server. May cause the database server to become unavailable temporarily.",
	Run:   home_backup_func,
}

func home_backup_func(cmd *cobra.Command, args []string) {
	if err := home_backup(cmd, args); err != nil {
		printwarn(err)
	}
}

func home_backup(cmd *cobra.Command, args []string) (err error) {
	wiz := wizard.New()

	err = wiz.Init()
	defer wiz.Close()

	var (
		requested_path string
		actual_path    string
	)

	requested_path, err = cmd.Flags().GetString("output-file")
	if err != nil {
		return
	}

	actual_path, err = wiz.TakeHomeBackup(requested_path)
	if err != nil {
		return
	}

	printok("Saved home backup file to", actual_path)
	return
}
