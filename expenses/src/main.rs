mod cli;
mod expense_handler;

use clap::Parser;
use cli::{Cli, Command};
use expense_handler::{Expense, ExpenseImpl, ExpenseList, ExpenseListImpl};

fn main() {
    let mut expense_data = ExpenseList::load();
    let args = Cli::parse();

    match args.command {
        Command::Add {
            description,
            amount,
        } => {
            let next_id: usize = expense_data.next_id;
            let expense: Expense = Expense::new(next_id, description, amount);
            expense_data.list.insert(next_id, expense);
            expense_data.next_id += 1;
            expense_data.save();
        }
        Command::Delete { id } => {
            expense_data.remove_item(id);
            expense_data.save();
        }
        Command::List { format } => {
            let expenses = expense_data.list.values();

            // TODO: Find a way to wrap a word lenght is greather than a certain number
            if let Some(format) = format {
                match format.as_str() {
                    "csv" => {
                        println!("id,description,amount,created_at");
                        for expense in expenses {
                            println!(
                                "{},{},{},{}",
                                expense.id, expense.description, expense.amount, expense.date
                            );
                        }
                    }
                    _ => eprintln!("Only 'csv' is available as option"),
                };
            } else {
                println!(" {}", "-".repeat(62));
                println!(
                    "| {0: <10} | {1: <20} | {2: <10}  | {3: <10} |",
                    "id", "description", "amount", "created_at"
                );
                println!(" {}", "-".repeat(62));
                for expense in expenses {
                    println!(
                        "| {0: <10} | {1: <20} | ${2: <10} | {3: <10} |",
                        expense.id, expense.description, expense.amount, expense.date
                    );
                }
                println!(" {}", "-".repeat(62));
            }
        }
        Command::Summary { month } => {
            if let Some(month) = month {
                match month {
                    1..=12 => {
                        // TODO: implement a filter by year and month, then apply
                        // a reduce to sum the amounts.
                        todo!()
                    }
                    _ => eprintln!("No valid month provided"),
                }
            } else {
                let expenses = expense_data
                    .list
                    .values()
                    .filter_map(|entry| Some(entry.amount))
                    .reduce(|a, b| a + b)
                    .unwrap_or(0.0);
                println!("Total expenses: ${}", expenses)
            }
        }
        Command::Update {
            id,
            description,
            amount,
        } => {
            let expense = expense_data.list.get_mut(&id).unwrap();
            expense.update(description, amount);
            expense_data.save();
        }
    }
}
