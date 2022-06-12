package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
add - adds a new task to our list
list - lists all of our incomplete tasks
do - marks a task as complete
*/
var version = "0.0.1"

var root_cmd = &cobra.Command{
	Use:     "taskmanager",
	Version: version,
	Aliases: []string{"tm"},
	Short:   "task manager - a simple CLI to managing tasks",
	Long:    ` Add-List-Do command are available for now.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := root_cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing CLI '%s'", err)
		os.Exit(1)
	}
}
