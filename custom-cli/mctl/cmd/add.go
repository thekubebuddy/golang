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
	Short: "add operation with n integer or float",
	Long: `add operation`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")
		addInt(args)
	},
}

func init() {
	operationCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addInt(args []string){
	var sum int
	for _, ival := range args{
		temp,err := strconv.Atoi(ival)
		if err !=nil{
			fmt.Println(err)
			os.Exit(1)
		}
		// fmt.Println(ival)
		sum = sum + temp
	}
	fmt.Printf("Sum of integers %s is %d",args,sum)
}

// for _, fval := range args {
// 	ftemp, err := strconv.ParseFloat(fval, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	sum = sum + ftemp
// }