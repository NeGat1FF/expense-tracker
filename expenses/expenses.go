package expenses

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	ID          int
	Date        time.Time
	Description string
	Amount      int
}

func (e Expense) String() string {
	year, month, day := e.Date.Date()
	return fmt.Sprintf("%-4d %-12s %-15s $%d", e.ID, fmt.Sprintf("%d-%2.2d-%2.2d", year, month, day), e.Description, e.Amount)
}

func ReadExpenses() ([]Expense, error) {
	expenses := make([]Expense, 0)
	data, err := os.ReadFile("expenses.json")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return expenses, err
	}

	if !errors.Is(err, os.ErrNotExist) {
		err = json.Unmarshal(data, &expenses)
		if err != nil {
			return expenses, err
		}
	}

	return expenses, nil
}

func DeleteExpense(id int) error {
	if id < 1 {
		return fmt.Errorf("id can't be less than one")
	}

	expenses, err := ReadExpenses()
	if err != nil {
		return err
	}

	found := false
	for i, exp := range expenses {
		if exp.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("expense with id %d not found", id)
	}

	return WriteExpenses(expenses)
}

func FilterExpenses(month int, expenses []Expense) []Expense {
	res := make([]Expense, 0)

	for _, exp := range expenses {
		if int(exp.Date.Month()) == month {
			res = append(res, exp)
		}
	}

	return res
}

func GetSummary(expenses []Expense) int {
	sum := 0
	for _, exp := range expenses {
		sum += exp.Amount
	}

	return sum
}

func WriteExpenses(expenses []Expense) error {
	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddExpense(desc string, amount int) (int, error) {
	expenses, err := ReadExpenses()
	if err != nil {
		return 0, err
	}

	newExpense := Expense{}
	if len(expenses) == 0 {
		newExpense.ID = 1
	} else {
		newExpense.ID = expenses[len(expenses)-1].ID + 1
	}
	newExpense.Date = time.Now()
	newExpense.Description = desc
	newExpense.Amount = amount

	expenses = append(expenses, newExpense)

	return newExpense.ID, WriteExpenses(expenses)
}
