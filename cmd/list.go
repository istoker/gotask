package cmd

import (
	"fmt"

	"github.com/istoker/gotask/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks",
	Long:  `Lists all the tasks and their status.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := task.GetTaskList()
		if err != nil {
			fmt.Println("Tasklist could not be retrieved. Aborting.")
			return
		}
		for index := 0; index < len(tasks); index++ {
			if tasks[index].Completed {
				fmt.Print("[x] ")
			} else {
				fmt.Print("[ ] ")
			}
			fmt.Printf("%d %s", tasks[index].Id, tasks[index].Text)
			fmt.Print("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
