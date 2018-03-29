package cmd

import (
	"fmt"

	"github.com/istoker/gotask/internal"
	"github.com/spf13/cobra"
)

func generateId(list []task.Task) int {
	if cap(list) == 0 {
		return 1
	}
	return list[len(list)-1].Id + 1
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new \"task description\" [\"more descriptions\"]",
	Short: "Adds new task(s) to the tasklist",
	Long: `Adds new task(s) to the tasklist with the provided descriptions.
	The date of creation will be added.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		list, err := task.GetTaskList()
		if err != nil {
			fmt.Println("Tasklist could not be retrieved. Aborting.")
			return
		}
		for i := range args {
			id := generateId(list)
			list = append(list, task.Task{args[i], false, id})
		}
		err = task.SaveTaskList(list)
		if err != nil {
			fmt.Println("Error while saving the tasklist. Please try again.")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
