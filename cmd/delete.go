/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NeGat1FF/expense-tracker/expenses"
	"github.com/spf13/cobra"
)

var id int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense entry by specifying its ID.",
	Long: `Use the delete command to remove an expense from your tracking system by providing the expense's unique ID. This command allows you to clean up your records by removing any incorrect or outdated entries. 
For example: 
	delete --id 1 
will delete the expense with ID 1.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := expenses.DeleteExpense(id)
		if err != nil {
			return err
		}

		fmt.Println("Expense deleted successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.Flags().IntVar(&id, "id", 0, "ID of expense you want to delete")
	deleteCmd.MarkFlagRequired("id")
}
