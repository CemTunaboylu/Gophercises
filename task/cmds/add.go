package cmds

import (
	"fmt"
	"task_manager/tasks"

	"github.com/spf13/cobra"
)

// add command for cobra
var add_cmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Adds the task into the task list",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := tasks.Add(args[0])
		fmt.Println(res)
	},
}

func init() {
	root_cmd.AddCommand(add_cmd)
}
