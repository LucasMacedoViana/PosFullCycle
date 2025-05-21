/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//name, _ := cmd.Flags().GetString("name")
		fmt.Println("category called: " + category)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("exists: ", exists)
		id, _ := cmd.Flags().GetInt16("id")
		fmt.Println("id: ", id)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("chamado antes do run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("chamado depois do run")
	},
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	return fmt.Errorf("ocorreu um erro")
	//},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	//categoryCmd.PersistentFlags().StringP("name", "n", "y", "Name of the category")
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "category name")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if the category exists")
	categoryCmd.PersistentFlags().Int16P("id", "i", 0, "ID of the category")

}
