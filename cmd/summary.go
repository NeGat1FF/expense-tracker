/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NeGat1FF/expense-tracker/expenses"
	"github.com/spf13/cobra"
)

var month int

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "View a summary of your total expenses, with an option to filter by month.",
	Long: `Use the summary command to get an overview of your total expenses. You can choose to view all expenses or filter the summary by a specific month. This allows you to track your spending more effectively over time. 
For example: 
	summary --month 9 
will show the total expenses for September.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		exp, err := expenses.ReadExpenses()
		if err != nil {
			return err
		}

		if month != 0 {
			exp = expenses.FilterExpenses(month, exp)
		}

		fmt.Printf("Total expenses: $%d", expenses.GetSummary(exp))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "Specify the month you want to see summary for")
}
