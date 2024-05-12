package wizard

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/Gophercraft/core/wizard"
	"github.com/spf13/cobra"
)

var session_cmd = &cobra.Command{
	Use:   "session",
	Short: "Create an alpha session file by creating a ticket from the server",
	Run:   session_func,
}

func session_func(cmd *cobra.Command, args []string) {
	if err := session(cmd, args); err != nil {
		printwarn(err)
	}
}

func session(cmd *cobra.Command, args []string) (err error) {
	wiz := wizard.New()

	err = wiz.Init()

	if errors.Is(err, wizard.ErrDisconnected) {
		fmt.Println(err)
		err = nil
	} else if err != nil {
		return
	}

	if !wiz.Connected() {
		return fmt.Errorf("to get a session ticket, you must first be connected to one\nUsage: gophercraft connect [home server address]")
	}

	// Look for executable
	var (
		wd                string
		directory_entries []fs.DirEntry
		found_alpha       = false
	)
	wd, err = os.Getwd()
	if err != nil {
		return
	}

	directory_entries, err = os.ReadDir(wd)
	if err != nil {
		return
	}

	for _, entry := range directory_entries {
		if entry.Name() == "WoWClient.exe" {
			found_alpha = true
			break
		}
	}

	if !found_alpha {
		printwarn("You do not appear to be inside an Alpha game directory.\nTo log in to an Alpha realm, you may want to copy wow.ses into that game directory.")
	}

	var ticket string
	ticket, err = wiz.GenerateLoginTicket()
	if err != nil {
		return
	}

	printok("Got ticket", ticket, "from server!")

	if err = os.WriteFile("wow.ses", []byte(ticket), 0700); err != nil {
		return
	}

	printok("Wrote ticket to session file!")

	// Close the wizard
	wiz.Close()
	return
}
