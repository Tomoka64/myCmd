package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(divideCmd)
}

var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Calculator of division.",
	Long:  "Calculator to perform the division.",
	Run: func(cmd *cobra.Command, args []string) {
		var n1 int
		var n2 int
		n1, _ = strconv.Atoi(args[0])
		n2, _ = strconv.Atoi(args[1])
		fmt.Println(n1 / n2)
	},
}
