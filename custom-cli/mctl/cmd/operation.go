/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// operationCmd represents the operation command
var operationCmd = &cobra.Command{
	Use:   "operation",
	Short: "match operation",
	Long:  `math operation support for add, substract, divide`,
}

func init() {
	rootCmd.AddCommand(operationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// operationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// operationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
