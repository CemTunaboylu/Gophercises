package cmds

import (
	"fmt"

	"task_manager/tasks"

	"github.com/spf13/cobra"
)

// list command for cobra
var comp_cmd = &cobra.Command{
	Use:     "comp",
	Aliases: []string{"co"},
	Short:   "Lists tasks that are completed",
	Run: func(cmd *cobra.Command, args []string) {
		res := tasks.Completed()
		fmt.Println(res)
	},
}

// adding all the other commands to the root command
func init() {
	root_cmd.AddCommand(comp_cmd)
}
