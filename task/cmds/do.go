package cmds

import (
	"fmt"
	"task_manager/tasks"

	"github.com/spf13/cobra"
)

// do command for cobra
var do_cmd = &cobra.Command{
	Use:   "do",
	Short: "Does/marks the given task as finished, and removes from the tasklist ",
	Args:  cobra.RangeArgs(1, 10),
	Run: func(cmd *cobra.Command, args []string) {
		res := tasks.Do(args...)
		fmt.Println(res)
	},
}

func init() {
	root_cmd.AddCommand(do_cmd)
}
