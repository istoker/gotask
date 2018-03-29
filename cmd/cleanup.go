package cmd

import (
	"fmt"
	"strconv"

	"github.com/istoker/gotask/internal"
	"github.com/spf13/cobra"
)

//var all bool

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove tasks from the list",
	Long:  `Remove tasks from the list. Default is removing the completed tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		list, err := task.GetTaskList()
		if err != nil {
			fmt.Println("Tasklist could not be retrieved. Aborting.")
			return
		}
		if cmd.Flag("delete").Value.String() != "-1" {
			for i := 0; i < len(list); i++ {
				id, _ := strconv.Atoi(cmd.Flag("delete").Value.String())
				if list[i].Id == id {
					list = append(list[:i], list[i+1:]...)
				}
			}
		} else if all, err := cmd.Flags().GetBool("all"); !all {
			if err != nil {
				fmt.Println("No valid all command given. Aborting.")
				return
			}
			for i := 0; i < len(list); i++ {
				if list[i].Completed {
					list = append(list[:i], list[i+1:]...)
				}
			}
		} else {
			list = nil
		}
		task.SaveTaskList(list)
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
	//cleanupCmd.Flags().BoolVarP(&all, "all", "a", false, "Cleanup all tasks (even uncompleted ones)")
	cleanupCmd.Flags().BoolP("all", "a", false, "Cleanup all tasks (even uncompleted ones)")
	cleanupCmd.Flags().IntP("delete", "d", -1, "Remove task with ID from the list.")

}
