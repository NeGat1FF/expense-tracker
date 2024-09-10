/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NeGat1FF/expense-tracker/expenses"
	"github.com/spf13/cobra"
)

var description string
var amount int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense entry with a description and amount.",
	Long: `Use the add command to record a new expense in your tracking system. 
You can provide a brief description of the expense and specify the amount of money spent. This command will store the details and assign a unique ID to each entry for easy reference. 
For example: 
	add --description 'Dinner' --amount 20. 
This will add an expense entry with 'Dinner' as the description and $20 as the amount.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := expenses.AddExpense(description, amount)
		if err != nil {
			return err
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", id)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of your expense")

	addCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Amount of money you spent")
	addCmd.MarkFlagRequired("amount")
}
