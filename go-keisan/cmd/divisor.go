package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(divisorCmd)
}

var divisorCmd = &cobra.Command{
	Use:   "divisor",
	Short: "Calculator of division.",
	Long:  "Calculator to perform the division.",
	Run: func(cmd *cobra.Command, args []string) {
		a := args[0]
		num, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}
		i64 := int64(num)
		fmt.Println(getDivisor(i64))
	},
}

func getDivisor(targetNum int64) []int64 {
	var retFrontSlice []int64
	var retBackSlice []int64
	var loopMax int64 = targetNum
	var i int64 = 2
	retFrontSlice = append(retFrontSlice, 1)
	if targetNum > 1 {
		retBackSlice = append(retBackSlice, targetNum)
	}
	for ; i < loopMax; i++ {
		if targetNum%i == 0 {
			loopMax = targetNum / i
			retFrontSlice = append(retFrontSlice, i)
			if i != loopMax {
				retBackSlice = append(retBackSlice, loopMax)
			}
		}
	}
	for j := len(retBackSlice) - 1; j >= 0; j-- {
		retFrontSlice = append(retFrontSlice, retBackSlice[j])
	}
	return retFrontSlice
}
