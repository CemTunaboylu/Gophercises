package cmds

import (
	"fmt"

	"task_manager/tasks"

	"github.com/spf13/cobra"
)

// length of the output of the list
var list_length int
var verbose bool

// list command for cobra
var list_cmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists tasks",
	Run: func(cmd *cobra.Command, args []string) {
		res := tasks.List(list_length, verbose)
		fmt.Println(res)
	},
}

// adding all the other commands to the root command
func init() {
	list_cmd.Flags().IntVarP(&list_length, "length", "l", 5, "Output a list of length 'len'")
	list_cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "List view is verbose")
	root_cmd.AddCommand(list_cmd)
}
