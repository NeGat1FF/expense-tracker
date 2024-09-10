# Expense Tracker CLI

A simple command-line tool for managing your personal expenses, written in Go using the Cobra library. This tool allows you to add, list, summarize, and delete expenses with ease.

## Features

- **Add an Expense**: Record a new expense with a description and amount.
- **List Expenses**: Display all recorded expenses with ID, date, description, and amount.
- **Delete an Expense**: Remove an expense by specifying its ID.
- **Summarize Expenses**: View the total expenses, optionally filtered by a specific month.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/expense-tracker-cli.git
   cd expense-tracker-cli
   ```

2. Build the project:
   ```bash
   go build -o expense-tracker
   ```

3. Run the CLI tool:
   ```bash
   ./expense-tracker
   ```

## Usage

### Add a New Expense

To add a new expense, use the `add` command with a description and amount:

```bash
./expense-tracker add --description "Dinner" --amount 20
```

### List All Expenses

To list all recorded expenses, use the `list` command:

```bash
./expense-tracker list
```

### Delete an Expense

To delete an expense by its ID, use the `delete` command:

```bash
./expense-tracker delete --id 1
```

### Summarize Expenses

To get a summary of all expenses or filter by a specific month, use the `summary` command:

- **View Total Expenses**:
  ```bash
  ./expense-tracker summary
  ```

- **View Expenses for a Specific Month**:
  ```bash
  ./expense-tracker summary --month 9
  ```

## Flags and Options

- `--description, -d`: Description of the expense (used with the `add` command).
- `--amount, -a`: Amount of the expense (used with the `add` command).
- `--id`: ID of the expense (used with the `delete` command).
- `--month, -m`: Month number for filtering expenses (used with the `summary` command).

https://roadmap.sh/projects/expense-tracker