# expense-tracker

Source: [expense-tracker](https://roadmap.sh/projects/expense-tracker)

## Commands

```sh
add <description> <amount>  # Add a record to the list
list                        # Show a table with all records
list --format <csv>         # (Optional) displays records as csv
summary                     # Show the total of all the records
summary --month <month>     # (Optional) Show the total of the records for a month in the current year
delete --id <id>            # Delete a record 
update --id <id> --description <description> --amount <amount>          # Update a record given an id, description and amount are optional
```