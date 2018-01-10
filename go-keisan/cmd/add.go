package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Calculator of addition.",
	Long:  "Calculator to perform the addition.",
	Run: func(cmd *cobra.Command, args []string) {
		sum := 0
		for _, arg := range args {
			var a int
			a, _ = strconv.Atoi(arg)
			sum += a
		}
		fmt.Println(sum)
	},
}
