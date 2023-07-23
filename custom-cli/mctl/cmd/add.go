/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add operation with 'n' positive integer or float",
	Long:  `add operation`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")
		// Conditioning based on the flag status
		fstatus, _ := cmd.Flags().GetBool("float")
		// if fstatus is true then call addFloat function
		if fstatus {
			addFloat(args)
		} else {
			addInt(args)
		}

	},
}

func init() {
	operationCmd.AddCommand(addCmd)
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//  Local flag for adding floating support
	addCmd.Flags().BoolP("float", "t", false, "floating addition support")
}

func addInt(args []string) {
	var sum int
	for _, ival := range args {
		temp, err := strconv.Atoi(ival)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// fmt.Println(ival)
		sum = sum + temp
	}
	fmt.Printf("Sum of integers %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, fval := range args {
		temp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// fmt.Println(fval)
		sum = sum + temp
	}
	fmt.Printf("Sum of floats %s is %.3f", args, sum)
}
