package cmd

import (
	"fmt"
	"strconv"

	"github.com/istoker/gotask/internal"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete the task with the given ID",
	Long:  `Complete the task with the given ID`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		list, err := task.GetTaskList()
		if err != nil {
			fmt.Println("Tasklist could not be retrieved. Aborting.")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Given parameter is not a valid ID. Aborting.")
		}
		found := false
		for i := 0; i < len(list); i++ {
			if list[i].Id == int(id) {
				list[i].Completed = true
				found = true
				break
			}
		}
		if !found {
			fmt.Println("The provided ID was not found in the task list")
		}
		task.SaveTaskList(list)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
