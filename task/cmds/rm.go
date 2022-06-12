package cmds

import (
	"fmt"
	"task_manager/tasks"

	"github.com/spf13/cobra"
)

var all bool

// rm command cobra
var rm_cmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Removes the specified task from the list",
	Args:    cobra.RangeArgs(1, 10),
	Run: func(cmd *cobra.Command, args []string) {
		res := tasks.Remove(all, args...)
		fmt.Println(res)
	},
}

func init() {
	rm_cmd.Flags().BoolVar(&all, "all", false, "Remove all tasks equals to clearing ")
	root_cmd.AddCommand(rm_cmd)
}
